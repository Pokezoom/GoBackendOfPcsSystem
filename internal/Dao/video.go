package Dao

import (
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/internal/mode"
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

// 获取视频页
func (d VideoData) GetVideoList(req mode.VideoListReq) ([]tables.Video, error) {
	var videos []tables.Video
	query := d.g.GDB().Table("video").Where("deleted = ?", 0)

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Class != "" {
		query = query.Where("class = ?", req.Class)
	}
	if req.Subject != "" {
		query = query.Where("subject = ?", req.Subject)
	}
	if req.StartDate != "" && req.EndDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}

	if err := query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
