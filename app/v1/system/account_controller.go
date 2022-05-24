// Package system
// @projectName DreamBlog
// @description 登陆|注册控制器
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月08日 18:05:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	"dream-blog/global"
	"dream-blog/model/common/response"
	systemReq "dream-blog/model/system/request"
	systemResp "dream-blog/model/system/response"
	"dream-blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login
// @Summary 用户登陆
// @Tags Account
// @Produce application/json
// @Description: 用户登陆
// @receiver b
// @param data body systemReq.Login true "用户名, 密码, 邮箱"
// @Success 200 {object} response.Response{data=systemResp.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /user/login [post]
//
func (b *BaseApi) Login(ctx *gin.Context) {
	var l systemReq.Login
	err := ctx.ShouldBindJSON(&l)
	if err != nil {
		return
	}
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) || !global.SYS_CONFIG.Captcha.Enable {
		if err, user := userService.Login(l); err != nil {
			global.SYS_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", ctx)
		} else {
			// 签发token
			b.signNext(ctx, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", ctx)
	}

}

// signNext
// @Description: token 签发 🛠️待增加多点登陆token的处理
//
func (b *BaseApi) signNext(ctx *gin.Context, user systemResp.RespUser) {
	j := &utils.JWT{SigningKey: []byte(global.SYS_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Username: user.Username,
		NickName: user.Nickname,
		Email:    user.Email,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.SYS_LOG.Error("获取token失败", zap.Error(err))
		response.FailWithMessage("获取token失败", ctx)
		return
	}
	response.OkWithDetailed(systemResp.LoginResponse{
		AccessToken:  token,
		RefreshToken: user.RefreshToken,
		ExpiresAt:    claims.ExpiresAt.Unix()},
		"登陆成功", ctx)
}

// Register
// @Tags Account
// @Summary 用户注册
// @Security ApiKeyAuth
// @Produce application/json
// @Description: 用户注册接口 /base/admin_register || /user/register admin_register需要进行鉴权
// @Receiver b
// @Param data body systemReq.Register true "用户名, 密码, 邮箱, 昵称 /base/admin_register || /user/register admin_register需要进行鉴权"
// @Success 200 {object} response.Response{msg=string} "账号注册成功信息"
// @Router /user/register [post]
//
func (b *BaseApi) Register(ctx *gin.Context) {
	var r systemReq.Register
	_ = ctx.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err, userReturn := userService.Register(r)
	if err != nil {
		global.SYS_LOG.Error("注册失败", zap.Error(err))
		response.FailWithDetailed(systemResp.SysUserResponse{User: userReturn}, "注册失败", ctx)
	} else {
		response.Success("注册成功", ctx)
	}
}

// ChangePassword
// @Tags Account
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {object} response.Response{msg=string} "用户修改密码"
// @Router /user/changePassword [post]
//
func (b *BaseApi) ChangePassword(ctx *gin.Context) {
	var user systemReq.ChangePasswordStruct
	_ = ctx.ShouldBindJSON(&user)
	// 验证是否符合要求
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err := userService.ChangePassword(&user)
	if err != nil {
		global.SYS_LOG.Error("密码修改失败", zap.Error(err))
		response.FailWithMessage("原密码错误！", ctx)
	} else {
		response.Success("修改成功", ctx)
	}
}
