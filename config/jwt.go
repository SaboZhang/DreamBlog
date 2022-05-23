// Package config
// @projectName DreamBlog
// @description jwt认证登录token生成
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月03日 20:25:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey  string `mapstructure:"signingKey" json:"signingKey" yaml:"signingKey"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expiresTime" json:"expiresTime" yaml:"expiresTime"` // 过期时间
	BufferTime  int64  `mapstructure:"bufferTime" json:"bufferTime" yaml:"bufferTime"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                // 签发者
}

type Claims struct {
	UserId   uint64
	Email    string
	UserName string
	NickName string
	jwt.RegisteredClaims
}
