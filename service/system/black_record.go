// Package system
// @projectName DreamBlog
// @description JWT黑名单实现
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 12:55:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	"context"
	"dream-blog/global"
	"dream-blog/model/system"
	"go.uber.org/zap"
	"time"
)

type JwtService struct{}

// LoadAll
// @Description: 从数据库中加载所有的JWT黑名单并且加入到BlackCache
//
func LoadAll() {
	var data []string
	err := global.SYS_DB.Model(&system.BlackRecord{}).Select("jti").Find(&data).Error
	if err != nil {
		global.SYS_LOG.Error("jwt黑名单加载失败", zap.Error(err))
		return
	}
	// jwt黑名单加入到 BlackCache 中
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct {
		}{})
	}

}

// IsBlackList
// @Description: 判断JWT是否属于黑名单
// @Receiver j
// @Param jwt string
// @return bool
//
func (jwtService *JwtService) IsBlackList(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

// GetRedisJWT
// @Description: 从Redis中获取JWT
// @Receiver jwtService
// @Param userName string
// @return err
// @return redisJWT
//
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.SYS_REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

// JonsInBlackList
// @Description: 拉黑jwt
// @Receiver jwtService
// @Param jwtList system.BlackRecord
// @return err
//
func (jwtService *JwtService) JonsInBlackList(jwtList system.BlackRecord) (err error) {
	err = global.SYS_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jti, struct{}{})
	return
}

// SetRedisJWT
// @Description: Set RedisJWT
// @Receiver jwtService
// @Param jwt string
// @Param userName string
// @return err
//
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.SYS_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.SYS_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
