/*
 * @Author: sucy suchunyu1998@gmail.com
 * @Date: 2023-11-17 20:34:27
 * @LastEditTime: 2024-03-10 00:06:05
 * @LastEditors: Suchunyu
 * @Description:
 * @FilePath: \GoBackendOfPcsSystem\main.go
 * Copyright (c) 2023 by Suchunyu, All Rights Reserved.
 */

package main

import (
	"GoDockerBuild/config"
	"GoDockerBuild/router"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	config.Test()
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

// package main

// import (
// 	"GoDockerBuild/internal/utils" // 根据你的项目结构调整导入路径
// 	"fmt"
// )

// func main() {
// 	err := utils.TestDBConnection()
// 	if err != nil {
// 		fmt.Println("Failed to connect to database:", err)
// 		return
// 	}
// 	fmt.Println("Connected to database successfully.")
// }
