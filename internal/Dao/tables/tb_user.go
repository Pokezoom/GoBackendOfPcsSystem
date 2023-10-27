package tables

import (
	"time"
)

type User struct {
	UserID      int       `gorm:"primaryKey;autoIncrement;comment:用户ID，主键"`
	Name        string    `gorm:"type:varchar(255);not null;unique;comment:用户名，唯一索引"`
	Password    string    `gorm:"type:varchar(255);not null;comment:md5后的用户密码"`
	Email       string    `gorm:"type:varchar(255);comment:电子邮件地址"`
	PhoneNumber string    `gorm:"type:varchar(20);comment:手机号码"`
	UserType    string    `gorm:"type:enum('1','2','3');not null;comment:用户类型（1-老师，2-学生，3-管理员）"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;comment:更新时间"`
	Deleted     bool      `gorm:"default:0;comment:是否删除（1为删除，0为存在）"`
}

func (u User) TableName() string {
	return "user"
}
