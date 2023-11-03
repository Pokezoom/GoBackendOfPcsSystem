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
