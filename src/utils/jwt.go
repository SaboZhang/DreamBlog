// Package utils
// @projectName DreamBlog
// @description JWT认证工具类
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月05日 00:38:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package utils

import (
	"crypto/rand"
	"dream-blog/global"
	"dream-blog/model/system/request"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token已过期")
	TokenNotValidYet = errors.New("token无效")
	TokenMalformed   = errors.New("非法token")
	TokenInvalid     = errors.New("无法处理此token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.SYS_CONFIG.JWT.SigningKey),
	}
}

// CreateClaims
// @Description: 创建Claims
// @receiver j
// @return request.CustomClaims
//
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	exp := time.Now().Unix() + global.SYS_CONFIG.JWT.ExpiresTime
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.SYS_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Local()), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),  // 过期时间 7天  配置文件
			Issuer:    global.SYS_CONFIG.JWT.Issuer,           // 签名的发行者
		},
	}
	return claims
}

// CreateToken
// @Description: 创建一个token
// @receiver j
// @return string
// @return error
//
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken
// @Description: 旧token 换新token 使用归并回源避免并发问题
// @receiver j
// @return string
// @return error
//
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.SysConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken
// @Description: 解析token
// @receiver j
// @return *request.CustomClaims
// @return error
//
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// token 是过期的
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

func AddRefreshToken() string {
	refreshToken := generateReferToken()
	return refreshToken
}

func generateReferToken() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
