package service

import (
	"GoDockerBuild/internal/Dao"
	"GoDockerBuild/middleware"
)

var Project ProjectService

func init() {
	middleware.Register(
		func() {
			Project = ProjectService{data: Dao.NewProjectData()}
		})
}

type ProjectService struct {
	data Dao.ProjectData
}

func (s ProjectService) CreateProject(name string) error {
	return s.data.Create(name)
}
