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

}
