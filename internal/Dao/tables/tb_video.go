package tables

import (
	"time"
)

type Video struct {
	ID           int       `gorm:"primaryKey;autoIncrement;comment:主键ID"`
	Name         string    `gorm:"type:varchar(255);not null;unique;comment:视频名称"`
	UserID       int       `gorm:"comment:上传用户ID"`
	Duration     int       `gorm:"not null;comment:视频总时长（秒）"`
	URL          string    `gorm:"type:varchar(512);not null;comment:视频URL"`
	Class        string    `gorm:"type:varchar(50);comment:所属班级"`
	AcademicYear string    `gorm:"type:varchar(50);comment:学年"`
	Subject      string    `gorm:"type:varchar(50);comment:科目"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP;comment:更新时间"`
	Deleted      bool      `gorm:"default:0;comment:是否删除（1为删除，0为存在）"`
}

func (i Video) TableName() string {
	return "video"
}
