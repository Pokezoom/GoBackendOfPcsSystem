package service

import (
	"GoDockerBuild/config"
	"GoDockerBuild/internal/Dao"
	"GoDockerBuild/internal/Dao/tables"
	"GoDockerBuild/internal/mode"
	"GoDockerBuild/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var Video VideoService

func init() {
	middleware.Register(
		func() {
			Video = VideoService{data: Dao.NewVideoData()}
		})
}

type VideoService struct {
	data Dao.VideoData
}

// 通过ID删除视频（软删除）
func (s VideoService) DeleteVideoById(videoId int, userId int) error {
	video, err := s.data.GetVideoById(videoId)
	if err != nil {
		return err
	}
	if video.UserID != userId {
		logrus.Infof("%d用户删除视频%d失败，原因没有权限", userId, videoId)
		return errors.New("该用户没有删除视频的权限")
	}
	logrus.Infof("%d用户成功删除了视频%d", userId, videoId)
	return s.data.DeleteVideoById(videoId)
}

// UploadAndSaveVideo 上传并保存视频
func (s VideoService) UploadAndSaveVideo(c *gin.Context, req mode.UploadReq) (int, error) {
	file, err := c.FormFile("video")
	if err != nil {
		return 0, err
	}
	if req.VideoName == "" {
		return 0, errors.New("缺少视频名称")
	}
	filesName := config.GetVideoPath()
	// 默认保存文件到桌面的某个文件夹里面
	desktopPath := filepath.Join(os.Getenv("HOME"), "Desktop", filesName)
	err = os.MkdirAll(desktopPath, os.ModePerm)
	if err != nil {
		return 0, err
	}

	filePath := filepath.Join(desktopPath, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return 0, err
	}

	// 创建视频记录
	video := tables.Video{
		Name:         req.VideoName,
		UserID:       req.UserID,
		URL:          filePath,
		Class:        req.Class,
		Duration:     req.Duration,
		Subject:      req.Subject,
		AcademicYear: req.AcademicYear,
	}

	err = s.data.CreateVideo(video)
	if err != nil {
		return 0, err
	}
	return video.ID, nil
}

func (s VideoService) GetVideoList(req mode.VideoListReq) ([]tables.Video, error) {
	return s.data.GetVideoList(req)
}

func (s VideoService) GetVideoPathByID(videoID int) (string, error) {
	video, err := s.data.GetVideoById(videoID)
	if err != nil {
		return "", err
	}
	return video.URL, nil
}
