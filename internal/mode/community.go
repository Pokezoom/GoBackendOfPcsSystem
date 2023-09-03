package mode

import (
	"time"
)

type Community struct {
	AppID                  string     `gorm:"primaryKey;type:varchar(64);not null;comment:应用id"`
	ID                     string     `gorm:"primaryKey;type:varchar(64);not null;comment:社群id"`
	Title                  string     `gorm:"type:varchar(64);not null;comment:社群名称"`
	Describe               *string    `gorm:"type:varchar(2048);comment:社群简介"`
	ImgURL                 *string    `gorm:"type:varchar(256);comment:社群配图url"`
	ImgURLCompressed       *string    `gorm:"type:varchar(256);comment:社群压缩配图url"`
	FeedsCount             *int       `gorm:"default:0;comment:动态数量"`
	UserCount              int        `gorm:"default:0;not null;comment:参与社群人数"`
	PaymentType            int        `gorm:"not null;comment:付费类型：1-免费、2-单笔、3-付费产品包"`
	PiecePrice             *int       `gorm:"comment:(字段已废弃,价格在统一商品表t_sku_x)-payment_type为2时，单笔价格（分）;payment_type为3时，专栏价格（分）"`
	StartAt                *time.Time `gorm:"type:timestamp;comment:上架时间"`
	StopAt                 *time.Time `gorm:"type:timestamp;comment:下架时间"`
	CommunityState         int        `gorm:"default:0;not null;comment:状态(0-上架 1-下架 2-删除,3-已解散)"`
	IsFeedsPush            int8       `gorm:"type:tinyint(2);default:0;not null;comment:是否开启群主动态推送：0-默认关闭，1-开启"`
	IsCommentPush          int8       `gorm:"type:tinyint(2);default:0;not null;comment:是否开启点赞和评论推送：0-默认关闭，1-开启"`
	CreateClient           int        `gorm:"default:0;not null;comment:默认0-PC管理台；1-手机端管理台 2-鹅社群小程序 3-鹅课程关联创建"`
	Creator                string     `gorm:"type:varchar(32);default:'';not null;comment:创建者"`
	CreatedAt              time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdatedAt              time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null;autoUpdateTime;comment:更新时间，有修改自动更新"`
	RightType              int8       `gorm:"default:0;not null;comment:社群权益类型 0-不限制加入 10-企业内所有员工 20-指定部门、成员加入"`
	CorpID                 string     `gorm:"type:varchar(128);default:'';not null;comment:授权方企业微信id、当为企业微信小程序创建的时候存在"`
	BarJSON                *string    `gorm:"type:varchar(4096);comment:导航栏排序数据(json)"`
	TagsJSON               *string    `gorm:"type:varchar(4096);comment:官方标签排序数据(json)"`
	MemberExpireType       int8       `gorm:"type:tinyint(2);default:0;not null;comment:会员有效期类型（0-永久，1-按年付费）"`
	MemberExpireTime       int        `gorm:"default:0;not null;comment:有效期（单位秒）"`
	MemberServiceSafeState int8       `gorm:"type:tinyint(2);default:0;not null;comment:是否开启服务保障（0-无，1-开启）"`
	MemberServiceSafeDays  int8       `gorm:"type:tinyint(2);default:0;not null;comment:服务保障天数（默认3天）"`
	OpenPrivateMessage     int8       `gorm:"type:tinyint(2);default:0;not null;comment:允许成员私信(0否1是)"`
	OpenQuestions          int8       `gorm:"type:tinyint(2);default:1;not null;comment:是否开启问答功能(0-否,1-是)"`
	RenewalDiscount        int        `gorm:"default:0;not null;comment:是否开启续费折扣 0--否 1--是"`
	RenewalNoticeDay       uint8      `gorm:"default:0;not null;comment:续费提醒天数"`
	Discount               float32    `gorm:"default:1;not null;comment:折扣比例"`
	DetailMongoID          *string    `gorm:"type:varchar(40);comment:圈子详情页链接,存储在mongo中"`
	JoinInvalidAt          *time.Time `gorm:"type:datetime;comment:加入截止时间"`
	TabsJSON               string     `gorm:"type:varchar(180);default:'';not null;comment:圈子底部tab列表json"`
}

// TableName sets the insert table name for this struct type
func (c *Community) TableName() string {
	return "t_community"
}
