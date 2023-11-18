package service

import (
	"GoDockerBuild/config"
	"GoDockerBuild/internal/Dao"
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/internal/mode"
	"GoDockerBuild/middleware"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/sync/errgroup"
	"gorm.io/datatypes"
	"os"
	"path/filepath"
	"strings"
)

var VideoAnalysis VideoAnalysisService

func init() {
	middleware.Register(
		func() {
			VideoAnalysis = VideoAnalysisService{data: Dao.NewVideoAnalysisData(), videoData: Dao.NewVideoData()}
		})
}

type VideoAnalysisService struct {
	data      Dao.VideoAnalysisData
	videoData Dao.VideoData
}

// 通过ID删除视频分析（软删除）
func (s VideoAnalysisService) DeleteVideoAnalysisById(ctx context.Context, id int) error {
	return s.data.DeleteVideoAnalysisById(id)
}

// 生成视频分析数据
func (s VideoAnalysisService) AnalysisVideo(ctx *gin.Context, req mode.VideoAnalysisReq) (tables.VideoAnalysis, error) {
	var res tables.VideoAnalysis
	video, err := s.videoData.GetVideoById(req.VideoId)
	if err != nil || video.ID == 0 || video.URL != "" {
		return res, err
	}
	// 创建一个errgroup
	G, ctx2 := errgroup.WithContext(ctx)

	// 根据请求字段调用相应接口
	if req.FacialData == 1 {
		G.Go(func() error {
			data, err := s.FacialData(ctx2)
			if err == nil {
				res.FacialData = data
			}
			return err
		})
	}
	if req.FatigueData == 1 {
		G.Go(func() error {
			data, err := s.FatigueData(ctx2)
			if err == nil {
				res.FatigueData = data
			}
			return err
		})
	}
	if req.LimbData == 1 {
		G.Go(func() error {
			data, err := s.LimbData(ctx2)
			if err == nil {
				res.LimbData = data
			}
			return err
		})
	}
	// 等待所有goroutine完成
	if err := G.Wait(); err != nil {
		return res, err
	}
	res.VideoID = video.ID
	res.UploaderID = req.UserID
	res.VideoURL = video.URL
	res.ImageURL = "" //图片展示用，这个后续要做
	// 存数据
	err = s.data.CreateVideoAnalysis(res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// 学生表情
func (s VideoAnalysisService) FacialData(ctx context.Context) (datatypes.JSON, error) {
	postBody := []byte{}
	url := config.GetAIUrl()
	FacialData, err := middleware.SendHTTPRequest("POST", url+"facialData", postBody)
	if err != nil {
		return nil, err
	}
	return FacialData, nil
}

// 疲劳数据
func (s VideoAnalysisService) FatigueData(ctx context.Context) (datatypes.JSON, error) {
	postBody := []byte{}
	url := config.GetAIUrl()
	Data, err := middleware.SendHTTPRequest("POST", url+"/fatigueData", postBody)
	if err != nil {
		return nil, err
	}
	return Data, nil
}

// 肢体数据
func (s VideoAnalysisService) LimbData(ctx context.Context) (datatypes.JSON, error) {
	postBody := []byte{}
	url := config.GetAIUrl()
	Data, err := middleware.SendHTTPRequest("POST", url+"/limbData", postBody)
	if err != nil {
		return nil, err
	}
	return Data, nil
}

// 生成视频分析的PDF报告
func (s VideoAnalysisService) GeneratePDFReport(ctx *gin.Context, req mode.GenerateReport) error {
	analysis, err := s.data.GetVideoAnalysisById(req.VideoAnalysisId)
	if err != nil {
		return err
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// 设置标题
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "视频分析报告", "0", 1, "C", false, 0, "")
	pdf.Ln(10) // 添加额外的空行

	// 添加数据到PDF的辅助函数
	addDataToPDF := func(title string, data string) {
		if data == "" {
			data = "N/A" // 处理空值
		}
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(0, 10, fmt.Sprintf("%s: %s", title, data), "0", 1, "", false, 0, "")
	}

	// 获取并添加视频信息
	video, err := s.videoData.GetVideoById(analysis.VideoID)
	if err != nil {
		return err
	}

	addDataToPDF("视频名称", video.Name)
	addDataToPDF("视频时长", fmt.Sprintf("%d 秒", video.Duration))
	addDataToPDF("所属班级", video.Class)
	addDataToPDF("学年", video.AcademicYear)
	addDataToPDF("科目", video.Subject)
	// 添加其他视频相关信息...

	pdf.Ln(5) // 在视频信息和分析数据之间添加额外的空行

	// 将 VideoAnalysis 字段添加到 PDF
	addDataToPDF("视频分析ID", fmt.Sprintf("%d", analysis.ID))
	addDataToPDF("创建人 ID", fmt.Sprintf("%d", analysis.UploaderID))
	addDataToPDF("视频 URL", analysis.VideoURL)
	addDataToPDF("图片 URL", analysis.ImageURL)
	// 为其他字段重复上面的步骤...
	report := config.GetReport()
	// 构建文件保存路径
	desktopPath := filepath.Join(os.Getenv("HOME"), "Desktop", report)
	safeFileName := strings.ReplaceAll(video.Name, " ", "_") + ".pdf" // 确保文件名是安全的
	pdfFilePath := filepath.Join(desktopPath, safeFileName)

	// 保存 PDF 文件
	return pdf.OutputFileAndClose(pdfFilePath)
}

// GetReportFilePath 返回给定报告的文件路径
func (s VideoAnalysisService) GetReportFilePath(reportName string) (string, error) {
	reportFolder := filepath.Join(os.Getenv("HOME"), "Desktop", "report")
	reportFilePath := filepath.Join(reportFolder, reportName)

	// 检查文件是否存在
	if _, err := os.Stat(reportFilePath); os.IsNotExist(err) {
		return "", errors.New("报告文件不存在")
	}

	return reportFilePath, nil
}

// VideoAnalysisService 是您的服务层结构体
func (s VideoAnalysisService) GetVideoAnalysisList(req mode.VideoAnalysisListReq) ([]tables.VideoAnalysis, error) {
	// 逻辑代码，执行基于分页的查询并返回视频分析数据列表
	return s.data.GetVideoAnalysisList(req)
}
