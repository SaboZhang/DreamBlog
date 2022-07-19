// Package model
// @projectName DreamBlog
// @description 基础结构体
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月03日 14:09:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package model

import (
	"database/sql/driver"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserBaseModel struct {
	CreateTime   time.Time      `json:"create_time" gorm:"type:datetime;default:NOW();comment:创建时间"` // 创建时间
	UpdateTime   *time.Time     `json:"update_time" gorm:"type:datetime;comment:更新时间"`               // 更新时间
	IsDeleted    gorm.DeletedAt `json:"is_deleted" gorm:"type:bit(1)"`                               // 删除标记
	DeletedTime  *time.Time     `json:"deleted_time" gorm:"type:datetime"`                           // 删除时间
	CreateUserId uint64         `json:"create_user_id" gorm:"comment:创建人ID"`                         // 创建人ID
	UpdateUserId uint64         `json:"update_user_id"`                                              // 更新人ID
	DeleteUserId uint64         `json:"delete_user_id"`                                              // 删除人ID
}

type BaseUUID struct {
	ID uuid.UUID `gorm:"type:char(36);primarykey;comment:主键Id"`
}

type BaseId struct {
	ID uint64 `gorm:"primarykey;comment:主键Id"`
}

// BeforeCreate
// @Description: UUID创建钩子
// @receiver u
// @return err
//
func (u *BaseUUID) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type CusBoolean bool

func (b CusBoolean) Value() (driver.Value, error) {
	result := make([]byte, 1)
	if b {
		result[0] = byte(1)
	} else {
		result[0] = 0
	}
	return result, nil
}

func (b CusBoolean) Scan(v interface{}) error {
	bytes := v.([]byte)
	if bytes[0] == 0 {
		b = false
	} else {
		b = true
	}
	return nil
}
