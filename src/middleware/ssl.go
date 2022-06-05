// Package middleware
// @projectName DreamBlog
// @description https模式
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月29日 21:13:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package middleware

import (
	"dream-blog/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go.uber.org/zap"
)

func TlsMode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(ctx.Writer, ctx.Request)
		if err != nil {
			global.SYS_LOG.Error("SSL模式出现错误", zap.Error(err))
			fmt.Println(err)
			return
		}
		ctx.Next()
	}
}
