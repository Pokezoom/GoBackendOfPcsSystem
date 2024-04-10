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
	"gorm.io/datatypes"
	"os"
	"path/filepath"
	"strings"
	"time"
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
func (s VideoAnalysisService) AnalysisVideo(ctx *gin.Context, req mode.VideoAnalysisReq) (string, error) {
	//修改视频标签
	err := s.videoData.UpdateVideo(tables.Video{
		ID:           req.VideoId,
		Name:         req.Name,
		Subject:      req.Subject,
		Class:        req.Class,
		AcademicYear: req.AcademicYear,
	})
	if err != nil {
		return "", err
	}
	res := tables.VideoAnalysis{
		Name:         req.Name,
		UploaderID:   req.UserID,
		VideoID:      req.VideoId,
		Subject:      req.Subject,
		Class:        req.Class,
		AcademicYear: req.AcademicYear,
		CreatedAt:    time.Now(),
		Deleted:      false,
		ID:           0,
	}

	// 存数据
	err = s.data.CreateVideoAnalysis(res)
	if err != nil {
		return "err", err
	}
	return "ok", nil
	//var res tables.VideoAnalysis
	//video, err := s.videoData.GetVideoById(req.VideoId)
	//if err != nil || video.ID == 0 || video.URL == "" {
	//	return res, err
	//}
	//// 创建一个errgroup
	//G, ctx2 := errgroup.WithContext(ctx)
	//
	//// 根据请求字段调用相应接口
	//if req.FacialData == 1 {
	//	G.Go(func() error {
	//		data, err := s.FacialData(ctx2)
	//		if err == nil {
	//			res.FacialData = data
	//		}
	//		return err
	//	})
	//}
	//if req.FatigueData == 1 {
	//	G.Go(func() error {
	//		data, err := s.FatigueData(ctx2)
	//		if err == nil {
	//			res.FatigueData = data
	//		}
	//		return err
	//	})
	//}
	//if req.LimbData == 1 {
	//	G.Go(func() error {
	//		data, err := s.LimbData(ctx2)
	//		if err == nil {
	//			res.LimbData = data
	//		}
	//		return err
	//	})
	//}
	//// 等待所有goroutine完成,这里记得改    应该是err != nil ,现在为了调试所以这样写
	//if err := G.Wait(); err == nil {
	//	return res, err
	//}
	//res.VideoID = video.ID
	//res.UploaderID = req.UserID
	//res.VideoURL = video.URL
	//res.ImageURL = "" //图片展示用，这个后续要做

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

func (s VideoAnalysisService) GeneratePDFReport(ctx *gin.Context, req mode.GenerateReport) error {
	analysis, err := s.data.GetVideoAnalysisById(req.VideoAnalysisId)
	if err != nil {
		return err
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// Set title
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Video Analysis Report", "0", 1, "C", false, 0, "")
	pdf.Ln(10) // Add an extra empty line

	// Helper function to add data to PDF
	addDataToPDF := func(title string, data string) {
		if data == "" {
			data = "N/A" // Handle empty value
		}
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(0, 10, fmt.Sprintf("%s: %s", title, data), "0", 1, "", false, 0, "")
	}

	// Fetch and add video information
	video, err := s.videoData.GetVideoById(analysis.VideoID)
	if err != nil {
		return err
	}

	addDataToPDF("Video Name", video.Name)
	addDataToPDF("Video Duration", fmt.Sprintf("%d seconds", video.Duration))
	addDataToPDF("Class", video.Class)
	addDataToPDF("Academic Year", video.AcademicYear)
	addDataToPDF("Subject", video.Subject)
	// Add other video related information...

	pdf.Ln(5) // Add an extra empty line between video info and analysis data

	// Add VideoAnalysis fields to PDF
	addDataToPDF("Video Analysis ID", fmt.Sprintf("%d", analysis.ID))
	addDataToPDF("Creator ID", fmt.Sprintf("%d", analysis.UploaderID))
	addDataToPDF("Video URL", analysis.VideoURL)
	addDataToPDF("Image URL", analysis.ImageURL)
	// Repeat for other fields...
	report := config.GetReport()
	// Build file save path
	desktopPath := filepath.Join(os.Getenv("HOME"), "Desktop", report)
	safeFileName := strings.ReplaceAll(video.Name, " ", "_") + ".pdf" // Ensure filename is safe
	pdfFilePath := filepath.Join(desktopPath, safeFileName)

	// Check and create directory before saving PDF
	if _, err := os.Stat(desktopPath); os.IsNotExist(err) {
		// If the folder does not exist, create it
		err := os.MkdirAll(desktopPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Unable to create folder: %v", err)
		}
	}

	// Save the PDF file
	return pdf.OutputFileAndClose(pdfFilePath)
}

// GetReportFilePath 返回给定报告的文件路径
func (s VideoAnalysisService) GetReportFilePath(reportName string) (string, error) {
	reportFolder := filepath.Join(os.Getenv("HOME"), "Desktop", "reports")
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
