// Package global
// @projectName DreamBlog
// @description TODO
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月04日 21:36:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package global

import (
	"dream-blog/config"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	SYS_DB                *gorm.DB
	SYS_REDIS             *redis.Client
	SYS_CONFIG            config.Server
	SYS_VP                *viper.Viper
	SYS_LOG               *zap.Logger
	SysConcurrencyControl = &singleflight.Group{}
	BlackCache            local_cache.Cache
)
