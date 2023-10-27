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

func (d ProjectData) CreateProject(name string) error {
	s := tables.Name_inof{Name: name}
	err := d.g.GDB().Table("name_info").Create(&s).Error
	return err
}

func (d ProjectData) DeleteProject(id int) error {
	s := tables.Name_inof{Id: id}
	err := d.g.GDB().Table("name_info").Delete(&s).Error
	return err
}

func (d ProjectData) GetProjectById(id int) (tables.Name_inof, error) {
	s := tables.Name_inof{}
	err := d.g.GDB().Table("name_info").Where("id = ?", id).First(&s).Error
	return s, err
}
