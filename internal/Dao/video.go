package Dao

import (
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/middleware"
	"time"
)

type VideoData struct {
	g middleware.EGorm
}

func NewVideoData() VideoData {
	return VideoData{middleware.EGorm{"video"}}
}

// 创建视频
func (d VideoData) CreateVideo(video tables.Video) error {
	err := d.g.GDB().Table("video").Create(&video).Error
	return err
}

// 通过ID删除视频（软删除）
func (d VideoData) DeleteVideoById(id int) error {
	video := tables.Video{ID: id, Deleted: true, UpdatedAt: time.Now()}
	err := d.g.GDB().Table("video").Where("id = ?", id).Updates(&video).Error
	return err
}

// 通过ID获取视频
func (d VideoData) GetVideoById(id int) (tables.Video, error) {
	video := tables.Video{}
	err := d.g.GDB().Table("video").Where("id = ? AND deleted = ?", id, 0).First(&video).Error
	return video, err
}

// 更新视频信息
func (d VideoData) UpdateVideo(video tables.Video) error {
	video.UpdatedAt = time.Now()
	err := d.g.GDB().Table("video").Where("id = ?", video.ID).Updates(&video).Error
	return err
}
