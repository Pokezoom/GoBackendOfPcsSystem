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
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, middleware.Response{500, "error", nil})
		return
	}
	fmt.Print("start create!")
	err = service.Project.CreateProject(req.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, middleware.Response{500, "error", nil})
		return
	}
	context.JSON(http.StatusOK, middleware.Response{200, "", nil})

	return
}
func (p ProjectController) Delete(context *gin.Context) {
	req := mode.ProDel{}
	context.ShouldBindJSON(&req)
	err := service.Project.DeleteProject(req.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, middleware.Response{500, err.Error(), nil})
		return
	}
	context.JSON(http.StatusOK, middleware.Response{200, "", nil})

	return
}
