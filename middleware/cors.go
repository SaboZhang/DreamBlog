// Package middleware
// @projectName DreamBlog
// @description 跨域中间件配置
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月21日 22:37:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package middleware

import (
	"dream-blog/config"
	"dream-blog/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors
// @Description: 直接放行所有请求
// @return gin.HandlerFunc
//
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}

// CorsByRules
// @Description: 跨域配置白名单模式
// @return gin.HandlerFunc
//
func CorsByRules() gin.HandlerFunc {
	// 放行全部
	if global.SYS_CONFIG.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(ctx *gin.Context) {
		whitelist := checkCors(ctx.GetHeader("origin"))

		// 通过检查添加请求头
		if whitelist != nil {
			ctx.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			ctx.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			ctx.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			ctx.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				ctx.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		// 严格白名单模式且未通过检查，直接拒绝
		if whitelist == nil && global.SYS_CONFIG.Cors.Mode == "strict-whitelist" && !(ctx.Request.Method == "GET" && ctx.Request.URL.Path == "/health") {
			ctx.AbortWithStatus(http.StatusForbidden)
		} else {
			// 非严格白名单 无论检查是否通过均放行OPTIONS
			if ctx.Request.Method == "OPTIONS" {
				ctx.AbortWithStatus(http.StatusNoContent)
			}

		}
		// 处理请求
		ctx.Next()
	}
}

// checkCors
// @Description: 遍历获取白名单配置
// @return *config.CORSWhitelist
//
func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.SYS_CONFIG.Cors.Whitelist {
		// 遍历配置中的跨域头，进行匹配
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
