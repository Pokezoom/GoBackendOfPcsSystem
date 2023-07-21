package main

import (
	"GoDockerBuild/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	router.Router(r)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Debug("go服务已开启")
	// 3.监听端口，默认在8080
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
