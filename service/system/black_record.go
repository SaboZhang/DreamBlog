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
	"dream-blog/global"
	"dream-blog/model/system"
	"go.uber.org/zap"
)

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
