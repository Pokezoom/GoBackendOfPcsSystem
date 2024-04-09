package controller

import (
	"GoDockerBuild/internal/mode"
	"GoDockerBuild/internal/service"
	"GoDockerBuild/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
func (u UserController) RegisterUser(c *gin.Context) {
	var req mode.RegistrationReq
	//bs, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	fmt.Println("ZF create error occurred")
	//	return
	//}
	//fmt.Printf("ZF create body:%s \n", bs)
	//
	//err = json.Unmarshal(bs, &req)
	//if err != nil {
	//	fmt.Println("json unmarshal failed", err.Error())
	//	return
	//}
	//fmt.Printf("LGY create req:%+v \n", req)

	// 从表单数据中手动提取字段并填充到req结构体中
	req.UserID, _ = strconv.Atoi(c.PostForm("userId"))
	req.Username = c.PostForm("username")
	req.Password = c.PostForm("password")
	req.Email = c.PostForm("email")
	req.PhoneNumber = c.PostForm("phoneNumber")
	// req.UserType, _ = strconv.Atoi(c.PostForm("userType"))
	req.UserType = c.PostForm("userType")

	userID, err := service.User.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, middleware.Response{
		Code: 200,
		Msg:  "注册成功",
		Data: mode.RegistrationRes{UserID: userID, Status: "Success"},
	})
}

// LoginUser 用户登录
func (u UserController) LoginUser(c *gin.Context) {
	var req mode.LoginReq
	// extrace username and password
	//bs, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	fmt.Printf("ZF [login] read request failed:%s \n", err.Error())
	//	return
	//}
	//err = json.Unmarshal(bs, &req)
	//if err != nil {
	//	fmt.Printf("ZF [login] json unmarshal failed:%s \n", err.Error())
	//	return
	//}
	//fmt.Printf("ZF [login] request:%+v \n", req)

	// 从表单数据中提取用户名和密码
	req.Username = c.PostForm("username")
	req.Password = c.PostForm("password")

	// 确保用户名和密码非空
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, middleware.Response{
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
		c.JSON(http.StatusUnauthorized, middleware.Response{
			Code: http.StatusUnauthorized,
			Msg:  "登录失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 如果登录成功，返回用户ID、用户名和成功消息
	c.JSON(http.StatusOK, middleware.Response{
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
