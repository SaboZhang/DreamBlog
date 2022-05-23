// Package middleware
// @projectName DreamBlog
// @description JWT认证中间件
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月05日 23:37:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package middleware

import (
	"dream-blog/global"
	"dream-blog/model/common/response"
	"dream-blog/model/system"
	"dream-blog/service"
	"dream-blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("authorization")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登陆或非法访问", ctx)
			ctx.Abort()
			return
		}
		if jwtService.IsBlackList(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "令牌失效或在异地登陆", ctx)
			ctx.Abort()
			return
		}
		j := utils.NewJWT()
		// 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权过期", ctx)
				ctx.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			exp := time.Now().Unix() + global.SYS_CONFIG.JWT.ExpiresTime
			claims.ExpiresAt = jwt.NewNumericDate(time.Unix(exp, 0))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			ctx.Header("new-token", newToken)
			ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			if global.SYS_CONFIG.System.UseMultipoint {
				// 用户多点登陆的处理 需借助Redis 或者其他共享方式
				err, RedisJwtToken := jwtService.GetRedisJWT(claims.Username)
				if err != nil {
					global.SYS_LOG.Error("从Redis获取token失败！", zap.Error(err))
				} else {
					// 获取成功拉黑原来的token
					_ = jwtService.JonsInBlackList(system.BlackRecord{Jti: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}

		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
