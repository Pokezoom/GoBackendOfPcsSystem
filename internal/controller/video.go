package controller

import (
	"GoDockerBuild/internal/mode"
	"GoDockerBuild/internal/service"
	"GoDockerBuild/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
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
	req := mode.UploadReq{}
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, middleware.Response{500, "error", nil})
		return
	}
	//TODO 权限鉴定，参数校验等
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
