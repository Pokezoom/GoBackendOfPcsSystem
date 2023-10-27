package Dao

import (
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/middleware"
	"time"
)

type UserData struct {
	g middleware.EGorm
}

func NewUserData() UserData {
	return UserData{middleware.EGorm{"user"}}
}

// CreateUser  创建用户
func (d UserData) CreateUser(user tables.User) error {
	err := d.g.GDB().Table("user").Create(&user).Error
	return err
}

// DeleteUserById 通过ID删除用户（软删除）
// 使用Updates方法并传入一个部分填充的结构体时，只有该结构体中非零字段会被更新。
func (d UserData) DeleteUserById(id int) error {
	user := tables.User{UserID: id, Deleted: true, UpdatedAt: time.Now()}
	err := d.g.GDB().Table("user").Where("user_id = ?", id).Updates(&user).Error
	return err
}

// GetUserById 通过ID获取用户
func (d UserData) GetUserById(id int) (tables.User, error) {
	user := tables.User{}
	err := d.g.GDB().Table("user").Where("user_id = ? AND deleted = ?", id, 0).First(&user).Error
	return user, err
}

// UpdateUser 更新用户信息
func (d UserData) UpdateUser(user tables.User) error {
	user.UpdatedAt = time.Now()
	err := d.g.GDB().Table("user").Where("user_id = ? AND deleted = ?", user.UserID, 0).Updates(&user).Error
	return err
}
