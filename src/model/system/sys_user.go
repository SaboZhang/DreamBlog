// Package system
// @projectName DreamBlog
// @description 系统用户结构体
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 13:07:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	"dream-blog/model"
	"gorm.io/gorm"
	"time"
)

type SysUser struct {
	model.BaseId
	Username          string              `gorm:"type:varchar(24);not null;comment:用户名"`
	Nickname          string              `gorm:"type:varchar(24);comment:昵称"`
	Avatar            string              `gorm:"type:varchar(255);comment:用户默认生成图像，为null、头像url"`
	Email             string              `gorm:"type:varchar(100);comment:电子邮箱"`
	Active            string              `gorm:"type:varchar(10);comment:是否激活"`
	PhoneNumber       string              `gorm:"type:varchar(100);comment:手机号"`
	Introduction      string              `gorm:"type:varchar(100);comment:个人介绍"`
	BlogAddress       string              `gorm:"type:varchar(100);comment:个人主页"`
	LastLoginTime     *time.Time          `json:"last_login_time" gorm:"datetime;comment:最后一次登录时间"`
	RefreshToken      string              `gorm:"type:varchar(200);comment:JWT 登录，保存生成的随机token值。"`
	IsEmailConfirmed  model.CusBoolean    `gorm:"type:bit(1);not null"`
	PasswordResetCode string              `gorm:"type:varchar(255)"`
	Salt              string              `gorm:"type:varchar(100);comment:盐值"`
	AuthorityId       string              `json:"authorityId" gorm:"default:9999;comment:用户角色ID"`
	SysUserIdentity   SysUserIdentity     `gorm:"foreignKey:CreateUserId;references:CreateUserId"`
	Base              model.UserBaseModel `gorm:"embedded"`
}

type SysUserIdentity struct {
	model.BaseUUID
	IdentityType    string         `gorm:"type:varchar(20);comment:认证类型， Password，GitHub、QQ、WeiXin等"`
	Identifier      string         `gorm:"type:varchar(24);comment:认证者，例如 用户名,手机号，邮件等，"`
	Credential      string         `gorm:"type:varchar(80);comment:凭证，例如 密码,存OpenId、Id，同一IdentityType的OpenId的值是唯一的"`
	ExtraProperties string         `gorm:"type:varchar(255)"`
	CreateTime      time.Time      `json:"create_time" gorm:"type:datetime;default:NOW();comment:创建时间"` // 创建时间
	UpdateTime      *time.Time     `json:"update_time" gorm:"type:datetime;comment:更新时间"`               // 更新时间
	IsDeleted       gorm.DeletedAt `json:"is_deleted" gorm:"type:bit(1)"`                               // 删除标记
	DeletedTime     *time.Time     `json:"deleted_time" gorm:"type:datetime"`                           // 删除时间
	CreateUserId    uint64         `json:"create_user_id" gorm:"comment:创建人ID"`                         // 创建人ID
	UpdateUserId    uint64         `json:"update_user_id"`                                              // 更新人ID
	DeleteUserId    uint64         `json:"delete_user_id"`                                              // 删除人ID
}

// TableName
// @Description: 指定数据库表名
// @receiver SysUser
// @return string 数据库表名
//
func (SysUser) TableName() string {
	return "sys_user"
}

func (SysUserIdentity) TableName() string {
	return "sys_user_identity"
}
