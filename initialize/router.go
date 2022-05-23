// Package initialize
// @projectName DreamBlog
// @description 总路由初始化
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月05日 01:47:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package initialize

import (
	_ "dream-blog/docs"
	"dream-blog/global"
	"dream-blog/middleware"
	"dream-blog/routes"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := routes.RouterGroupApp.System

	Router.StaticFS(global.SYS_CONFIG.Local.Path, http.Dir(global.SYS_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.SYS_LOG.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.SYS_LOG.Info("register swagger handler")

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	PublicGroup.Use(middleware.RecaptchaVerify())
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	PrivateGroup.Use(middleware.RecaptchaVerify()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}
	global.SYS_LOG.Info("路由注册成功！")
	// 抛出指针
	return Router
}
