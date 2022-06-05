// Package utils
// @projectName DreamBlog
// @description Claims工具类
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 10:16:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package utils

import (
	"dream-blog/global"
	systemReq "dream-blog/model/system/request"
	"github.com/gin-gonic/gin"
)

func GetClaims(ctx *gin.Context) (*systemReq.CustomClaims, error) {
	token := ctx.Request.Header.Get("authorization")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.SYS_LOG.Error("从Gin的Context解析token失败")
	}
	return claims, err
}
