package Dao

import (
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/middleware"
)

type ProjectData struct {
	g middleware.EGorm
}

func NewProjectData() ProjectData {
	return ProjectData{middleware.EGorm{"project"}}
}

func (d ProjectData) Create(name string) error {
	s := tables.Name_inof{Name: name, Id: 12}
	err := d.g.GDB().Table("name_info").Create(&s).Error
	return err
}
