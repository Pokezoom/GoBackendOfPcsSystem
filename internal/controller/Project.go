package controller

import (
	"GoDockerBuild/internal/mode"
	"GoDockerBuild/internal/service"
	"GoDockerBuild/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Project ProjectController

type ProjectController struct {
}

func init() {
	middleware.Register(func() {
		Project = ProjectController{}
	})
}

func (p ProjectController) Create(context *gin.Context) {
	req := mode.ProCreate{}
	context.ShouldBindJSON(&req)
	fmt.Print("start create!")
	err := service.Project.CreateProject(req.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, middleware.Response{500, "error", nil})
		return
	}
	context.JSON(http.StatusOK, middleware.Response{200, "", nil})

	return
}
