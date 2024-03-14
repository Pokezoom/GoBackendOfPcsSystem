package controller

import (
	"GoDockerBuild/internal/mode"
	"GoDockerBuild/internal/service"
	"GoDockerBuild/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var User UserController

type UserController struct {
}

func init() {
	middleware.Register(func() {
		User = UserController{}
	})
}

// RegisterUser 用户注册
func (u UserController) RegisterUser(context *gin.Context) {
	var req mode.RegistrationReq

	// 从表单数据中手动提取字段并填充到req结构体中
	req.UserID, _ = strconv.Atoi(context.PostForm("userId"))
	req.Username = context.PostForm("username")
	req.Password = context.PostForm("password")
	req.Email = context.PostForm("email")
	req.PhoneNumber = context.PostForm("phoneNumber")
	// req.UserType, _ = strconv.Atoi(context.PostForm("userType"))
	req.UserType = context.PostForm("userType")

	userID, err := service.User.CreateUser(req)
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
		Msg:  "注册成功",
		Data: mode.RegistrationRes{UserID: userID, Status: "Success"},
	})
}

// LoginUser 用户登录
func (u UserController) LoginUser(context *gin.Context) {
    var req mode.LoginReq
    // 从表单数据中提取用户名和密码
    req.Username = context.PostForm("username")
    req.Password = context.PostForm("password")

    // 确保用户名和密码非空
    if req.Username == "" || req.Password == "" {
        context.JSON(http.StatusBadRequest, middleware.Response{
            Code: http.StatusBadRequest,
            Msg:  "用户名和密码不能为空",
            Data: nil,
        })
        return
    }

    // 调用登录服务，传入用户名和密码
    loginResponse, err := service.User.Login(req)
    if err != nil {
        // 如果服务返回错误，可能是因为用户名或密码不正确
        context.JSON(http.StatusUnauthorized, middleware.Response{
            Code: http.StatusUnauthorized,
            Msg:  "登录失败: " + err.Error(),
            Data: nil,
        })
        return
    }

    // 如果登录成功，返回用户ID、用户名和成功消息
    context.JSON(http.StatusOK, middleware.Response{
        Code: 200,
        Msg:  "登录成功",
        Data: loginResponse, // 直接使用 service 返回的响应对象
    })
}


// 	token, userID, err := service.User.Login(req)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, middleware.Response{
// 			Code: http.StatusBadRequest,
// 			Msg:  err.Error(),
// 			Data: nil,
// 		})
// 		return
// 	}

// 	context.JSON(http.StatusOK, middleware.Response{
// 		Code: 200,
// 		Msg:  "登录成功",
// 		Data: mode.LoginRes{UserID: userID, Token: token, Status: "Success"},
// 	})
// }

// // DeleteUser 删除用户
// func (u UserController) DeleteUser(context *gin.Context) {
// 	userID, err := strconv.Atoi(context.Param("userID"))
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, middleware.Response{
// 			Code: http.StatusBadRequest,
// 			Msg:  "无效的用户ID",
// 			Data: nil,
// 		})
// 		return
// 	}

// 	err = service.User.Delete(userID)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, middleware.Response{
// 			Code: http.StatusBadRequest,
// 			Msg:  err.Error(),
// 			Data: nil,
// 		})
// 		return
// 	}

// 	context.JSON(http.StatusOK, middleware.Response{
// 		Code: 200,
// 		Msg:  "删除成功",
// 		Data: nil,
// 	})
// }

// // UpdateUser 更新用户信息
// // 这个方法需要根据具体的业务逻辑来实现，这里只提供了一个大致的框架
// func (u UserController) UpdateUser(context *gin.Context) {
// 	var req mode.User
// 	if err := context.ShouldBindJSON(&req); err != nil {
// 		context.JSON(http.StatusBadRequest, middleware.Response{
// 			Code: http.StatusBadRequest,
// 			Msg:  err.Error(),
// 			Data: nil,
// 		})
// 		return
// 	}

// 	err := service.User.Update(req)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, middleware.Response{
// 			Code: http.StatusBadRequest,
// 			Msg:  err.Error(),
// 			Data: nil,
// 		})
// 		return
// 	}

// 	context.JSON(http.StatusOK, middleware.Response{
// 		Code: 200,
// 		Msg:  "更新成功",
// 		Data: nil,
// 	})
// }

// 其他用户相关的方法可以根据需要继续添加...
