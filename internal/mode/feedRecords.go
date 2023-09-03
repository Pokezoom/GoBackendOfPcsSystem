package mode

import (
	"time"
)

type CommunityFeedsRecord struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;comment:自增id"`
	AppID       string    `gorm:"type:varchar(64);not null;comment:动态所属店铺id"`
	CommunityID string    `gorm:"type:varchar(64);not null;comment:动态所属社群id"`
	FeedsID     string    `gorm:"type:varchar(64);not null;uniqueIndex:uidx_record;comment:动态或问答id"`
	UnionID     string    `gorm:"type:varchar(64);not null;uniqueIndex:uidx_record;comment:阅读的用户微信id/企微id"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null;autoUpdateTime;comment:更新时间"`
}

// TableName sets the insert table name for this struct type
func (c *CommunityFeedsRecord) TableName() string {
	return "t_community_feeds_record"
}
