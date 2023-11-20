/*
 * @Author: sucy suchunyu1998@gmail.com
 * @Date: 2023-11-19 16:02:29
 * @LastEditTime: 2023-11-19 16:02:30
 * @LastEditors: Suchunyu
 * @Description:
 * @FilePath: /GoBackendOfPcsSystem/router/router.go
 * Copyright (c) 2023 by Suchunyu, All Rights Reserved.
 */
package router

import (
	"GoDockerBuild/internal/controller"
	"GoDockerBuild/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Router 注册路由
func Router(r *gin.Engine) {
	initProject(r)

}
func initProject(r *gin.Engine) {
	r.Use(middleware.RecoveryMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	//ProjectRouter := r.Group("/project", middleware.AuthProject())
	//ProjectRouter.POST("", controller.Project.Create)
	//ProjectRouter.DELETE("", controller.Project.Delete)
	/*
		<———————————————视频相关的路由————————————————>
	*/
	videoRouter := r.Group("/video")
	videoRouter.POST("/upload", controller.Video.UploadVideo)
	videoRouter.DELETE("/delete", controller.Video.DelVideo)
	videoRouter.POST("/list", controller.Video.VideoList)                            //视频列表，支持模糊查询
	videoRouter.GET("/:videoID/download", controller.Video.DownloadVideo)            //下载视频
	videoRouter.GET("/:videoID/stream", controller.Video.PlayVideo)                  //实时播放视频
	videoRouter.POST("/analysis", controller.Video.AnalysisVideo)                    // 生成视频数据
	videoRouter.POST("/report/generate", controller.Video.GenerateReport)            // 生成视频数据pdf报告
	videoRouter.GET("/report/download/:reportName", controller.Video.DownloadReport) // 下载视频pdf报告
	videoRouter.POST("/analysis/list", controller.Video.VideoAnalysisList)

	/*
		<———————————————user相关的路由————————————————>
	*/

	logrus.Debug("路由注册完成")
}
