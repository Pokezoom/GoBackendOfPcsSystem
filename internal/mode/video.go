package mode

import (
	"gorm.io/datatypes"
	"time"
)

type UploadReq struct {
	UserID       int    `json:"userId"`
	VideoName    string `json:"videoName"`
	Class        string `json:"class"`
	AcademicYear string `json:"academicYear"`
	Subject      string `json:"subject"`
	Duration     int    `json:"duration"`
}
type UploadRes struct {
	VideoId int `json:"videoId"`
}
type DelVideoRes struct {
	UserID  int `json:"userId"`
	VideoId int `json:"videoId"`
}
type VideoListReq struct {
	PageSize  int    `json:"pageSize"`
	PageNum   int    `json:"pageNum"`
	Name      string `json:"name,omitempty"`
	Class     string `json:"class,omitempty"`
	Subject   string `json:"subject,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}
type VideoAnalysisReq struct {
	UserID          int `json:"userId"` //当前用户
	VideoId         int `json:"videoId"`
	FacialData      int `json:"facialData"` // 0-不需要，1-需要 以下相同
	FatigueData     int `json:"fatigueData"`
	LimbData        int `json:"limbData"`
	StudyStatusData int `json:"studyStatusData"`
}
type GenerateReport struct {
	UserID          int `json:"userId"` //当前用户
	VideoAnalysisId int `json:"videoAnalysisId"`
}

type VideoAnalysisListReq struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
}
type VideoAnalysisRes struct {
	ID                int            `json:"id" gorm:"primaryKey;autoIncrement;comment:分析ID，主键"`
	VideoID           int            `json:"videoId" gorm:"not null;comment:视频ID"`
	StudentAttendance int            `json:"studentAttendance" gorm:"not null;default:0;comment:学生出勤人数"`
	FacialData        datatypes.JSON `json:"facialData" gorm:"comment:表情数据，结构体"`
	FatigueData       datatypes.JSON `json:"fatigueData" gorm:"comment:疲劳数据，结构体"`
	LimbData          datatypes.JSON `json:"limbData" gorm:"comment:肢体数据，结构体"`
	StudyStatusData   datatypes.JSON `json:"studyStatusData" gorm:"comment:学习状态数据"`
	ImageURL          string         `json:"imageUrl" gorm:"type:varchar(512);not null;default:'';comment:图片URL"`
	VideoURL          string         `json:"videoUrl" gorm:"type:varchar(512);not null;default:'';comment:视频URL"`
	UploaderID        int            `json:"uploaderId" gorm:"comment:上传用户ID"`
	CreatedAt         time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt         time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP;comment:更新时间"`
	Deleted           bool           `json:"deleted" gorm:"default:0;comment:是否删除（1为删除，0为存在）"`
	Subject           string         `json:"subject" `
	Name              string         `json:"name" `
}
