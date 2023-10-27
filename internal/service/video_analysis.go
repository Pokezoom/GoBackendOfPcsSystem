package service

import (
	"GoDockerBuild/internal/Dao"
	"GoDockerBuild/middleware"
)

var VideoAnalysis VideoAnalysisService

func init() {
	middleware.Register(
		func() {
			VideoAnalysis = VideoAnalysisService{data: Dao.NewVideoAnalysisData()}
		})
}

type VideoAnalysisService struct {
	data Dao.VideoAnalysisData
}

// 通过ID删除视频分析（软删除）
func (s VideoAnalysisService) DeleteVideoAnalysisById(id int) error {
	return s.data.DeleteVideoAnalysisById(id)
}
