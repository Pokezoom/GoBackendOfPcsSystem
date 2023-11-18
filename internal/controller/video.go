package controller

import (
	"GoDockerBuild/internal/mode"
	"GoDockerBuild/internal/service"
	"GoDockerBuild/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var Video VideoController

type VideoController struct {
}

func init() {
	middleware.Register(func() {
		Video = VideoController{}
	})
}

func (V VideoController) UploadVideo(context *gin.Context) {
	// 初始化UploadReq结构体
	var req mode.UploadReq

	// 获取表单字段并填充到req结构体中
	req.UserID, _ = strconv.Atoi(context.PostForm("userId"))
	req.VideoName = context.PostForm("videoName")
	req.Class = context.PostForm("class")
	req.AcademicYear = context.PostForm("academicYear")
	req.Subject = context.PostForm("subject")
	req.Duration, _ = strconv.Atoi(context.PostForm("duration"))

	// TODO: 这里可以添加字段验证逻辑

	// 调用service层方法
	videoID, err := service.Video.UploadAndSaveVideo(context, req)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	context.JSON(http.StatusOK, middleware.Response{
		Code: 200,
		Msg:  "ok",
		Data: mode.UploadRes{videoID},
	})
}

func (V VideoController) DelVideo(context *gin.Context) {
	delReq := mode.DelVideoRes{}
	err := context.ShouldBindJSON(&delReq)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	videoId := delReq.VideoId
	userId := delReq.UserID
	err = service.Video.DeleteVideoById(videoId, userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	context.JSON(http.StatusOK, middleware.Response{
		Code: 200,
		Msg:  "ok",
		Data: mode.UploadRes{videoId},
	})
}

func (V VideoController) VideoList(context *gin.Context) {
	var req mode.VideoListReq
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	videos, err := service.Video.GetVideoList(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	context.JSON(http.StatusOK, middleware.Response{
		Code: 200,
		Msg:  "ok",
		Data: videos,
	})
}
func (V VideoController) DownloadVideo(context *gin.Context) {
	videoIDStr := context.Param("videoID")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  "无效的视频ID",
			Data: nil,
		})
		return
	}

	videoPath, err := service.Video.GetVideoPathByID(videoID)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	//可以通过访问 /play/:videoID 这个URL来播放视频
	context.File(videoPath)
}

func (V VideoController) PlayVideo(context *gin.Context) {
	videoIDStr := context.Param("videoID")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  "无效的视频ID",
			Data: nil,
		})
		return
	}

	videoPath, err := service.Video.GetVideoPathByID(videoID)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	//可以通过访问 /play/:videoID 这个URL来播放视频
	context.File(videoPath)
}

// AnalysisVideo 分析视频
func (V VideoController) AnalysisVideo(context *gin.Context) {
	var req mode.VideoAnalysisReq
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	res, err := service.VideoAnalysis.AnalysisVideo(context, req)
	if err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	context.JSON(http.StatusOK, middleware.Response{
		Code: 200,
		Msg:  "ok",
		Data: res,
	})
}

// GenerateReport 生成视频分析报告
func (v VideoController) GenerateReport(context *gin.Context) {
	var req mode.GenerateReport
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	err := service.VideoAnalysis.GeneratePDFReport(context, req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, middleware.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	context.JSON(http.StatusOK, middleware.Response{
		Code: 200,
		Msg:  "报告生成成功",
		Data: nil,
	})
}

// DownloadReport 用于下载视频分析报告
func (v VideoController) DownloadReport(context *gin.Context) {
	reportName := context.Param("reportName") // 假设报告名称通过 URL 参数传递

	filePath, err := service.VideoAnalysis.GetReportFilePath(reportName)
	if err != nil {
		context.JSON(http.StatusNotFound, middleware.Response{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	context.Header("Content-Disposition", "attachment; filename="+reportName)
	context.Header("Content-Type", "application/pdf")

	// 提供文件下载
	context.File(filePath)
}

func (V VideoController) VideoAnalysisList(context *gin.Context) {
	var req mode.VideoAnalysisListReq
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	analysisList, err := service.VideoAnalysis.GetVideoAnalysisList(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, middleware.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	context.JSON(http.StatusOK, middleware.Response{
		Code: 200,
		Msg:  "ok",
		Data: analysisList,
	})
}
