/*
 * @Author: Su Chunyu su18237805830@163.com
 * @Date: 2024-03-09 15:55:58
 * @LastEditTime: 2024-03-10 16:37:57
 * @LastEditors: Suchunyu
 * @Description:
 * @FilePath: \GoBackendOfPcsSystem\router\router.go
 * Copyright (c) 2024 by Suchunyu, All Rights Reserved.
 */
package router

import (
	"GoDockerBuild/internal/controller"
	"GoDockerBuild/middleware"
	"github.com/gin-contrib/cors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Router 注册路由
func Router(r *gin.Engine) {
	initProject(r)

}
func initProject(r *gin.Engine) {
	r.Use(middleware.RecoveryMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                                                            // 允许任何源
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},                                             // 允许的方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}, // 允许的头信息
		ExposeHeaders:    []string{"Content-Length"},                                                                               // 公开的头信息
		AllowCredentials: true,                                                                                                     // 凭证共享
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour, // 预检请求的结果缓存最大时长
	}))
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
	userRouter := r.Group("/user")
	userRouter.POST("/create", controller.User.RegisterUser) // 创建用户
	// userRouter.DELETE("/delete/:userID", controller.User.DeleteUser) // 删除用户
	// userRouter.GET("/list", controller.User.UserList)                // 用户列表
	userRouter.POST("/login", controller.User.LoginUser) // 用户登录
	// userRouter.PUT("/update/:userID", middleware.AuthUser(), controller.User.UpdateUser)       // 更新用户信息
	// userRouter.POST("/change-password", middleware.AuthUser(), controller.User.ChangePassword) // 修改密码
	// userRouter.POST("/reset-password", controller.User.ResetPassword)                          // 重置密码
	// userRouter.GET("/detail/:userID", middleware.AuthUser(), controller.User.UserDetail)       // 用户详情

	logrus.Debug("路由注册完成")

	logrus.Debug("路由注册完成")
}
