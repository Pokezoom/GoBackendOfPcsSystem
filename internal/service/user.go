/*
 * @Author: sucy suchunyu1998@gmail.com
 * @Date: 2023-11-17 21:06:46
 * @LastEditTime: 2023-11-17 21:06:48
 * @LastEditors: Suchunyu
 * @Description:
 * @FilePath: /GoBackendOfPcsSystem/internal/service/user.go
 * Copyright (c) 2023 by Suchunyu, All Rights Reserved.
 */
package service

import (
	"GoDockerBuild/internal/Dao"
	"GoDockerBuild/middleware"
)

var User UserService

func init() {
	middleware.Register(
		func() {
			User = UserService{data: Dao.NewUserData()}

		})
}

type UserService struct {
	data Dao.UserData
}

// // 通过ID删除用户（软删除）
// func (s UserService) DeleteUserById(id int) error {
// 	return s.data.DeleteUserById(id)
// }
