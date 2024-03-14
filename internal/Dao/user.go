/*
 * @Author: Su Chunyu su18237805830@163.com
 * @Date: 2024-03-09 15:55:58
 * @LastEditTime: 2024-03-10 17:35:09
 * @LastEditors: Suchunyu
 * @Description:
 * @FilePath: \GoBackendOfPcsSystem\internal\Dao\user.go
 * Copyright (c) 2024 by Suchunyu, All Rights Reserved.
 */
package Dao

import (
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/internal/mode"

	// "GoDockerBuild/internal/utils"
	"GoDockerBuild/middleware"
	// 确保导入路径根据实际项目结构调整
)

type UserData struct {
	g middleware.EGorm
}

func NewUserData() UserData {
	return UserData{middleware.EGorm{"user"}}
}

// CreateUser 创建用户
func (d UserData) CreateUser(user tables.User) (int, error) {
	// 使用GORM的Create方法创建用户
	// 注意：Create操作后，user对象将被填充包括ID在内的所有字段
	err := d.g.GDB().Table("user").Create(&user).Error
	if err != nil {
		return 0, err // 创建失败，返回错误
	}
	// 创建成功，返回用户ID
	return user.UserID, nil
}

func (ud *UserData) GetUserByUsername(username string) (*mode.User, error) {
	var user mode.User
	// 假设 d.g.GDB() 返回一个 *gorm.DB 实例
	err := ud.g.GDB().Table("user").Where("NAME = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// // CheckUserNameAndPassword 根据用户名和密码查询一条记录
// func (ud *UserData) CheckUserNameAndPassword(username, password string) (*mode.User, error) {
// 	user := &mode.User{}
// 	err := utils.Db.QueryRow("select id, username, password, email from users where username=? and password=?", username, password).Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// // // CheckUserName 根据用户名查询一条记录
// // func (ud *UserData) CheckUserName(username string) (*mode.User, error) {
// // 	user := &mode.User{}
// // 	err := utils.Db.QueryRow("select id, username, password, email from users where username=?", username).Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return user, nil
// // }

// // GetUserById 根据ID查询一条记录
// func (ud *UserData) GetUserById(id int) (*mode.User, error) {
// 	user := &mode.User{}
// 	err := utils.Db.QueryRow("select id, username, password, email from users where id=?", id).Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// // DeleteUserById 通过ID软删除用户
// func (ud *UserData) DeleteUserById(id int) error {
// 	_, err := utils.Db.Exec("update users set deleted=1 where id=?", id)
// 	return err
// }

// // UpdateUser 更新用户信息
// func (ud *UserData) UpdateUser(user *mode.User) error {
// 	_, err := utils.Db.Exec("update users set username=?, password=?, email=? where id=?", user.UserName, user.Password, user.Email, user.ID)
// 	return err
// }
