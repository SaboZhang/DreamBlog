// Package response
// @projectName DreamBlog
// @description 用户响应数据
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 17:18:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package response

import (
	"dream-blog/model/system"
	"github.com/google/uuid"
)

type RespUser struct {
	UUID         uuid.UUID
	ID           uint64
	Username     string
	Nickname     string
	Avatar       string
	Email        string
	PhoneNumber  string
	Introduction string
	Credential   string
	RefreshToken string
}

type SysUserResponse struct {
	User system.SysUser
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expiresAt"`
}
