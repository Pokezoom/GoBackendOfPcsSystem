package controller

import (
	"GoDockerBuild/internal/mode"
	"GoDockerBuild/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
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
	context.ShouldBindJSON()
	fmt.Print("start create!")
}
