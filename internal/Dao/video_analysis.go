package Dao

import (
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/middleware"
	"time"
)

type VideoAnalysisData struct {
	g middleware.EGorm
}

func NewVideoAnalysisData() VideoAnalysisData {
	return VideoAnalysisData{middleware.EGorm{"video_analysis"}}
}

// 创建视频分析
func (d VideoAnalysisData) CreateVideoAnalysis(videoAnalysis tables.VideoAnalysis) error {
	err := d.g.GDB().Table("video_analysis").Create(&videoAnalysis).Error
	return err
}

// 通过ID删除视频分析（软删除）
func (d VideoAnalysisData) DeleteVideoAnalysisById(id int) error {
	videoAnalysis := tables.VideoAnalysis{ID: id, Deleted: true, UpdatedAt: time.Now()}
	err := d.g.GDB().Table("video_analysis").Where("id = ?", id).Updates(&videoAnalysis).Error
	return err
}

// 通过ID获取视频分析
func (d VideoAnalysisData) GetVideoAnalysisById(id int) (tables.VideoAnalysis, error) {
	videoAnalysis := tables.VideoAnalysis{}
	err := d.g.GDB().Table("video_analysis").Where("id = ? AND deleted = ?", id, 0).First(&videoAnalysis).Error
	return videoAnalysis, err
}

// 更新视频分析信息
func (d VideoAnalysisData) UpdateVideoAnalysis(videoAnalysis tables.VideoAnalysis) error {
	videoAnalysis.UpdatedAt = time.Now()
	err := d.g.GDB().Table("video_analysis").Where("id = ?", videoAnalysis.ID).Updates(&videoAnalysis).Error
	return err
}
