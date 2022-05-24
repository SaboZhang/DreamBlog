// Package system
// @projectName DreamBlog
// @description 用户认证总路由
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 10:59:00
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

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	//userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemAppGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)       // 用户注册路由
		userRouter.POST("changePassword", baseApi.ChangePassword) // 修改密码路由
	}
}
