// Package request
// @projectName DreamBlog
// @description JWT请求结构体
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月05日 00:49:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package request

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint64
	Username    string
	NickName    string
	Email       string
	AuthorityId string
}
