// Package initialize
// @projectName DreamBlog
// @description 初始化Redis
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月24日 22:17:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package initialize

import (
	"context"
	"dream-blog/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Redis
// @Description: 初始化Redis
//
func Redis() {
	redisCfg := global.SYS_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		DB:       redisCfg.DB,
		Password: redisCfg.Password,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.SYS_LOG.Error("redis 连接失败, err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("redis 连接成功！response: ", zap.String("pong", pong))
		global.SYS_REDIS = client
	}
}
