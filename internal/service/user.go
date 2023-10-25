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

// 通过ID删除用户（软删除）
func (s UserService) DeleteUserById(id int) error {
	return s.data.DeleteUserById(id)
}
