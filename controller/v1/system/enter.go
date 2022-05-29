// Package system
// @projectName DreamBlog
// @description 路由入口
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月08日 18:06:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import "dream-blog/service"

type AppGroup struct {
	BaseApi
}

var (
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
