/*
 * @Author: Su Chunyu su18237805830@163.com
 * @Date: 2024-03-09 15:55:58
 * @LastEditTime: 2024-03-10 17:43:34
 * @LastEditors: Suchunyu
 * @Description:
 * @FilePath: \GoBackendOfPcsSystem\internal\mode\user.go
 * Copyright (c) 2024 by Suchunyu, All Rights Reserved.
 */
package mode

import "time"

// User 结构体调整以匹配 tables.User 的定义
type User struct {
	UserID      int       `json:"userId" gorm:"column:user_id"`                     // 假设数据库列名为 user_id
	Name        string    `json:"username" gorm:"column:NAME"`                      // 假设数据库列名为 name
	Password    string    `json:"password" gorm:"column:PASSWORD"`                  // 假设数据库列名为 password
	Email       string    `json:"email" gorm:"column:email"`                        // 假设数据库列名为 email
	PhoneNumber string    `json:"phoneNumber,omitempty" gorm:"column:phone_number"` // 假设数据库列名为 phone_number，且为可选字段
	UserType    string    `json:"userType" gorm:"column:user_type"`                 // 假设数据库列名为 user_type
	CreatedAt   time.Time `json:"-" gorm:"column:created_at"`                       // 假设数据库列名为 created_at
	UpdatedAt   time.Time `json:"-" gorm:"column:updated_at"`                       // 假设数据库列名为 updated_at
	Deleted     bool      `json:"-" gorm:"column:deleted"`                          // 假设数据库列名为 deleted
}

// 用户注册请求，新增字段以匹配 User 结构
type RegistrationReq struct {
	UserID      int    `json:"userId"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber,omitempty"` // 新增可选手机号码字段
	UserType    string `json:"userType"`              // 新增用户类型字段
}

// 用户注册响应，保持不变
type RegistrationRes struct {
	UserID int    `json:"userId"`
	Status string `json:"status"`
}

// 用户登录请求，保持不变
type LoginReq struct {
	Username string `json:"username" gorm:"column:NAME"`    
	Password string `json:"password" gorm:"column:PASSWORD"`
}

// 用户登录响应，保持不变
type LoginRes struct {
	UserID  int    `json:"userId"`
	Username string `json:"username" gorm:"column:NAME"`    
	Token   string `json:"token"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
