/*
 * @Author: sucy suchunyu1998@gmail.com
 * @Date: 2023-11-10 18:52:51
 * @LastEditTime: 2023-11-17 21:03:52
 * @LastEditors: Suchunyu
 * @Description: 
 * @FilePath: /GoBackendOfPcsSystem/internal/mode/user.go
 * Copyright (c) 2023 by Suchunyu, All Rights Reserved. 
 */

package mode


type User struct {
	ID       int64
	UserName string
	Password string
	Email    string
}

// 用户注册请求
type RegistrationReq struct {
	ID   int    `json:"userId"`
    Username string `json:"username"`
    Password string `json:"password"`
    Email    string `json:"email"`
    FullName string `json:"fullName"`
}

// 用户注册响应
type RegistrationRes struct {
    UserID int    `json:"userId"`
    Status string `json:"status"`
}

// 用户登录请求
type LoginReq struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// 用户登录响应
type LoginRes struct {
    UserID  int    `json:"userId"`
    Token   string `json:"token"`
    Status  string `json:"status"`
    Message string `json:"message,omitempty"`
}

