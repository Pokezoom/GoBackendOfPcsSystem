/*
 * @Author: Su Chunyu su18237805830@163.com
 * @Date: 2024-03-09 15:55:58
 * @LastEditTime: 2024-03-10 17:49:10
 * @LastEditors: Suchunyu
 * @Description:
 * @FilePath: \GoBackendOfPcsSystem\internal\service\user.go
 * Copyright (c) 2024 by Suchunyu, All Rights Reserved.
 */
package service

import (
	Dao "GoDockerBuild/internal/Dao" // 确保包名正确，且全部小写
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/internal/mode"
	"errors"
	"fmt"
)

var User UserService

type UserService struct {
	data Dao.UserData // 确保类型名正确
}

func (s UserService) CreateUser(req mode.RegistrationReq) (int, error) {

	// 创建用户实例
	newUser := tables.User{
		UserID:      req.UserID,
		Name:        req.Username,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		UserType:    req.UserType,
	}

	fmt.Printf("newUser:%+v", newUser)
	// 使用UserData的CreateUser方法创建用户
	userID, err := s.data.CreateUser(newUser)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// // DeleteUser 删除用户
// func (s UserService) DeleteUser(id int) error {
// 	return s.data.DeleteUserById(id)
// }

// // GetUserList 获取用户列表
// func (s UserService) GetUserList() ([]mode.User, error) {
// 	// 此处应根据实际业务逻辑调整，可能需要过滤掉已删除的用户
// 	return nil, errors.New("未实现")
// }

// // Login 用户登录
func (s UserService) Login(req mode.LoginReq) (*mode.LoginRes, error) {
	// 根据用户名获取用户信息
	user, err := s.data.GetUserByUsername(req.Username)
	if err != nil {
		// 处理用户不存在的情况
		return nil, err
	}

	// 直接比较明文密码
	if user.Password != req.Password {
		// 密码不匹配
		return nil, errors.New("invalid username or password")
	}

	// 登录成功，返回用户ID和用户名
	return &mode.LoginRes{
		UserID:   user.UserID,
		Username: user.Name,
	}, nil
}

// // UpdateUser 更新用户信息
// func (s UserService) UpdateUser(user *mode.User) error {
// 	// 这里应当处理密码加密等逻辑
// 	return s.data.UpdateUser(user)
// }

// // GetUserDetail 获取用户详细信息
// func (s UserService) GetUserDetail(id int) (*mode.User, error) {
// 	return s.data.GetUserById(id)
// }

// // EncryptPassword 密码加密示例函数（假设实现）
// func EncryptPassword(password string) (string, error) {
// 	// 使用bcrypt或其他加密库
// 	return password, nil // 示例代码，应替换为实际加密逻辑
// }

// // CheckPassword 密码校验示例函数（假设实现）
// func CheckPassword(password, hashedPassword string) bool {
// 	// 使用bcrypt或其他库来比较密码和哈希值
// 	return true // 示例代码，应替换为实际比较逻辑
// }
