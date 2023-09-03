package models

import (
	"time"
)

type AppBaseInfo struct {
	AppID           string    `gorm:"type:varchar(64);primaryKey;not null;comment:店铺ID"`
	MerchantID      string    `gorm:"type:varchar(64);comment:商户id"`
	AppName         string    `gorm:"type:varchar(64);comment:店铺名称"`
	VersionType     int       `gorm:"comment:店铺类型"`
	LastVersionType int       `gorm:"default:0;not null;comment:上一个版本"`
	AuthenticState  int8      `gorm:"type:tinyint;default:0;comment:店铺认证状态"`
	AuthenticBody   string    `gorm:"type:varchar(128);comment:店铺认证主体信息"`
	SignTime        time.Time `gorm:"type:timestamp;default:'0000-00-00 00:00:00';not null;comment:注册时间"`
	ExpireTime      time.Time `gorm:"type:datetime;default:'0000-00-00 00:00:00';comment:版本过期时间"`
	OpenedTime      time.Time `gorm:"type:timestamp;default:'0000-00-00 00:00:00';not null;comment:开通时间"`
	VersionStartAt  time.Time `gorm:"type:timestamp;default:'0000-00-00 00:00:00';not null;comment:当前版本开始时间"`
	// ... (continue for other fields)

	// Timestamp fields
	LastLoginAt     time.Time `gorm:"type:timestamp;default:'0000-00-00 00:00:00';not null;comment:店铺最近活跃时间"`
	LastPayAt       time.Time `gorm:"type:timestamp;default:'0000-00-00 00:00:00';not null;comment:店铺最近成交时间"`
	TransTime       time.Time `gorm:"type:timestamp;default:'0000-00-00 00:00:00';not null;comment:企学院大客户转化时间"`
	PayTime         time.Time `gorm:"type:timestamp;default:'0000-00-00 00:00:00';not null;comment:购买版本时间"`
	RecentPayTime   string    `gorm:"type:varchar(64);default:'0000-00-00 00:00:00';comment:最近一次续费时间"`
	LastFollowAt    time.Time `gorm:"type:timestamp;default:'0000-00-00 00:00:00';not null;comment:店铺最近跟进时间"`
	IsKaV2          int8      `gorm:"type:tinyint;default:1;not null;comment:是否是KA店铺"`
	StaffName       string    `gorm:"type:varchar(255);comment:服务管家"`
	StaffDepartment string    `gorm:"type:varchar(255);comment:服务管家组别"`
	NewIndustryName string    `gorm:"type:varchar(255);comment:行业"`
}

// TableName sets the insert table name for this struct type
func (c *AppBaseInfo) TableName() string {
	return "t_app_base_info"
}
