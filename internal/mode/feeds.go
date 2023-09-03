package mode

import (
	"time"
)

type CommunityFeeds struct {
	AppID            string     `gorm:"type:varchar(64);primary_key;not null;column:app_id" json:"app_id"`
	ID               string     `gorm:"type:varchar(64);primary_key;not null;column:id" json:"id"`
	CommunityID      string     `gorm:"type:varchar(64);not null" json:"community_id"`
	Title            *string    `gorm:"type:varchar(128);null" json:"title"`
	UserID           string     `gorm:"type:varchar(64);not null" json:"user_id"`
	UniversalUnionID string     `gorm:"type:varchar(64);not null;default:''" json:"universal_union_id"`
	UserType         int        `gorm:"type:int;default:0" json:"user_type"`
	FeedsType        int        `gorm:"type:int;default:0" json:"feeds_type"`
	OrgContent       *string    `gorm:"type:mediumtext;null" json:"org_content"`
	Content          *string    `gorm:"type:mediumtext;null" json:"content"`
	TagsContent      string     `gorm:"type:mediumtext;not null" json:"tags_content"`
	FileURL          string     `gorm:"type:varchar(256);not null;default:''" json:"file_url"`
	FileName         string     `gorm:"type:varchar(256);not null;default:''" json:"file_name"`
	FileSize         float64    `gorm:"type:float;default:0" json:"file_size"`
	FileJSON         string     `gorm:"type:varchar(4096);default:''" json:"file_json"`
	FeedsState       int        `gorm:"type:int;default:0" json:"feeds_state"`
	IsChosen         int        `gorm:"type:int;default:0" json:"is_chosen"`
	IsNotice         int        `gorm:"type:int;default:0" json:"is_notice"`
	SetChosenAt      *time.Time `gorm:"type:timestamp;null" json:"set_chosen_at"`
	SetNoticeAt      *time.Time `gorm:"type:timestamp;null" json:"set_notice_at"`
	ZanNum           int        `gorm:"type:int;default:0" json:"zan_num"`
	CommentCount     int        `gorm:"type:int;default:0" json:"comment_count"`
	SendType         int        `gorm:"type:int;default:0" json:"send_type"`
	IsExercise       int        `gorm:"type:int(2);default:0" json:"is_exercise"`
	IsPublic         int        `gorm:"type:tinyint(2);default:1" json:"is_public"`
	ShareNum         int        `gorm:"type:int;default:0" json:"share_num"`
	IP               string     `gorm:"type:varchar(39);default:''" json:"ip"`
	PushState        int        `gorm:"type:int(4);default:0" json:"push_state"`
	PushStateOutline int        `gorm:"type:tinyint(2);default:0" json:"push_state_outline"`
	CreatedAt        time.Time  `gorm:"type:timestamp;not null;default:'0000-00-00 00:00:00'" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"type:datetime;null" json:"deleted_at"`
	ApplyType        string     `gorm:"type:varchar(64);not null;default:''" json:"apply_type"`
	ExtendField1     string     `gorm:"type:varchar(256);not null;default:''" json:"extend_field1"`
	ExtendField2     string     `gorm:"type:varchar(256);not null;default:''" json:"extend_field2"`
	ExtendField3     string     `gorm:"type:varchar(512);not null;default:''" json:"extend_field3"`
}

func (CommunityFeeds) TableName() string {
	return "t_community_feeds"
}
