// Package system
// @projectName DreamBlog
// @description 用户管理
// @author 张涛
// @version 1.0.0
// @createTime 2022年06月15日 08:56:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	"dream-blog/global"
	"dream-blog/model/common/response"
	"dream-blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetInformation
// @Tags 用户管理
// @Summary 获取单个用户信息
// @Security ApiKeyAuth
// @Produce application/json
// @Description: 获取单个用户或者当前登陆用户的信息
// @Receiver b
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取用户信息"
// @Router /user/information [get]
//
func (b *BaseApi) GetInformation(ctx *gin.Context) {
	uid := utils.GetUserID(ctx)
	err, reqUser := userService.GetInformation(uid)
	if err != nil {
		global.SYS_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(gin.H{"information": reqUser}, "获取成功", ctx)
	}
}
