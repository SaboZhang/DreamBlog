// Package initialize
// @projectName DreamBlog
// @description orm自动建表
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月04日 22:13:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package initialize

import (
	"dream-blog/global"
	"dream-blog/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// RegisterTables
// @Description: 自动注册数据库表
//
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.BlackRecord{},
		system.SysUser{},
		system.SysUserIdentity{},
	)
	if err != nil {
		global.SYS_LOG.Error("自动建表失败", zap.Error(err))
		os.Exit(0)
	}
	global.SYS_LOG.Info("自动建表成功")
}
