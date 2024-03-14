/*
 * @Author: Su Chunyu su18237805830@163.com
 * @Date: 2024-03-09 15:55:58
 * @LastEditTime: 2024-03-10 16:16:04
 * @LastEditors: Suchunyu
 * @Description:
 * @FilePath: \GoBackendOfPcsSystem\internal\Dao\tables\tb_user.go
 * Copyright (c) 2024 by Suchunyu, All Rights Reserved.
 */
package tables

import (
	"time"
)

type User struct {
	UserID      int       `json:"user_id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null;unique;comment:用户名，唯一索引"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null;comment:md5后的用户密码"`
	Email       string    `json:"email" gorm:"type:varchar(255);comment:电子邮件地址"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(20);comment:手机号码"`
	UserType    string    `json:"user_type" gorm:"not null;comment:用户类型（1-老师，2-学生，3-管理员）"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
	Deleted     bool      `json:"deleted" gorm:"default:0;comment:是否删除（1为删除，0为存在）"`
}

func (u User) TableName() string {
	return "user"
}
