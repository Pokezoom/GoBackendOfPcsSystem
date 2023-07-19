package main

import (
	"GoDockerBuild/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	router.Router(r)
	// 3.监听端口，默认在8080
	r.Run(":8000")
}

func Start() {

}
