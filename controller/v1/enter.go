// Package v1
// @projectName DreamBlog
// @description 路由入口集合
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月08日 18:05:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package v1

import "dream-blog/controller/v1/system"

type AppGroup struct {
	SystemAppGroup system.AppGroup
}

var ApiGroupApp = new(AppGroup)
