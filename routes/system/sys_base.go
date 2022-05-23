// Package system
// @projectName DreamBlog
// @description 基础路由
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月22日 22:07:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	v1 "dream-blog/app/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemAppGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)       // 登陆路由
		baseRouter.POST("captcha", baseApi.Captcha)   // 验证码路由
		baseRouter.POST("register", baseApi.Register) // 普通用户注册
	}
	return baseRouter
}
