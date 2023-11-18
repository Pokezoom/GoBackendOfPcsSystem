/*
 * @Author: sucy suchunyu1998@gmail.com
 * @Date: 2023-11-10 19:18:43
 * @LastEditTime: 2023-11-17 21:02:24
 * @LastEditors: Suchunyu
 * @Description: 
 * @FilePath: /GoBackendOfPcsSystem/internal/Dao/user.go
 * Copyright (c) 2023 by Suchunyu, All Rights Reserved. 
 */
package Dao

import (
	"GoDockerBuild/middleware"
	"GoDockerBuild/internal/utils"
	"GoDockerBuild/internal/mode"
)

type UserData struct {
	g middleware.EGorm
}

func NewUserData() UserData {
	return UserData{middleware.EGorm{"user"}}
}

// CreateUser  创建用户
func CreateUser(username string, password string, email string) error {
	sqlStr := "insert into users(username,password,email) values (?,?,?)"
	_, err := utils.Db.Exec(sqlStr, username, password, email)

	return err
}

// CheckUserName 根据用户名和密码查询一条记录
func CheckUserNameAndPassword(username string, password string) (*mode.User, error) {
	sqlStr := "select id,username,password,email from users where username=? and password=?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &mode.User{}
	row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
	return user, nil
}
// CheckUserName  根据用户名查询一条记录
func CheckUserName(username string) (*mode.User, error) {
	sqlStr := "select id,username,password,email from users where username=?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &mode.User{}
	row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
	return user, nil
}

// GetUserById 根据ID查询一条记录
func GetUserById(id int) (*mode.User, error) {
    sqlStr := "select id, username, password, email from users where id=?"
    row := utils.Db.QueryRow(sqlStr, id)
    user := &mode.User{}
    err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
    if err != nil {
        return nil, err
    }
    return user, nil
}





// DeleteUserById 通过ID软删除用户
func DeleteUserById(id int) error {
    sqlStr := "update users set deleted=1 where id=?"
    _, err := utils.Db.Exec(sqlStr, id)
    return err
}




// UpdateUser 更新用户信息
func UpdateUser(user *mode.User) error {
    sqlStr := "update users set username=?, password=?, email=? where id=?"
    _, err := utils.Db.Exec(sqlStr, user.UserName, user.Password, user.Email, user.ID)
    return err
}


