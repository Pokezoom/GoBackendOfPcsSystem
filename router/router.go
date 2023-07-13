package router

import (
	"GoDockerBuild/internal/controller"
	"GoDockerBuild/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router 注册路由
func Router(r *gin.Engine) {
	initProject(r)

}
func initProject(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	ProjectRouter := r.Group("/project", middleware.AuthProject())
	ProjectRouter.POST("", controller.Project.Create)
}
