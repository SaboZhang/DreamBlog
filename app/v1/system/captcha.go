// Package system
// @projectName DreamBlog
// @description 验证码控制器（api接口）
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 11:12:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	"dream-blog/global"
	"dream-blog/model/common/response"
	systemResp "dream-blog/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct {
}

// Captcha
// @Tags Account
// @Summary 获取验证码
// @Produce application/json
// @Accept application/json
// @Description: 获取验证码的base64图片
// @Success 200 {object} response.Response{data=systemResp.SysCaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router /base/captcha [post]
//
func (b *BaseApi) Captcha(ctx *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.SYS_CONFIG.Captcha.ImgHeight, global.SYS_CONFIG.Captcha.ImgWidth, global.SYS_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.SYS_LOG.Error("验证码获取失败", zap.Error(err))
		response.FailWithMessage("验证码获取失败！", ctx)
	} else {
		response.OkWithDetailed(systemResp.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.SYS_CONFIG.Captcha.KeyLong,
		}, "验证码获取成功", ctx)
	}
}
