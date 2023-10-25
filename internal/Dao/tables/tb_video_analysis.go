package tables

import (
	"gorm.io/datatypes"
	"time"
)

type VideoAnalysis struct {
	ID                int            `gorm:"primaryKey;autoIncrement;comment:分析ID，主键"`
	VideoID           int            `gorm:"not null;comment:视频ID"`
	StudentAttendance int            `gorm:"not null;default:0;comment:学生出勤人数"`
	FacialData        datatypes.JSON `gorm:"comment:表情数据，结构体"`
	FatigueData       datatypes.JSON `gorm:"comment:疲劳数据，结构体"`
	LimbData          datatypes.JSON `gorm:"comment:肢体数据，结构体"`
	StudyStatusData   datatypes.JSON `gorm:"comment:学习状态数据"`
	ImageURL          string         `gorm:"type:varchar(512);not null;default:'';comment:图片URL"`
	VideoURL          string         `gorm:"type:varchar(512);not null;default:'';comment:视频URL"`
	UploaderID        *int           `gorm:"comment:上传用户ID"`
	CreatedAt         time.Time      `gorm:"default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt         time.Time      `gorm:"default:CURRENT_TIMESTAMP;comment:更新时间"`
	Deleted           bool           `gorm:"default:0;comment:是否删除（1为删除，0为存在）"`
}

func (va VideoAnalysis) TableName() string {
	return "video_analysis"
}
