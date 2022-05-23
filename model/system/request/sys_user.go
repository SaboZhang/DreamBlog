// Package request
// @projectName DreamBlog
// @description 前端参数接收DTO
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月04日 11:57:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package request

type Login struct {
	Username  string `form:"username" json:"username" uri:"username" xml:"username"`                    // 用户名
	Password  string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"` // 密码
	Email     string `form:"email" json:"email" uri:"email" xml:"email"`                                // 邮箱
	Captcha   string `json:"captcha"`                                                                   // 验证码
	CaptchaId string `json:"captchaId"`                                                                 // 验证码ID
}

type Register struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"` // 用户名
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"` // 密码
	Email    string `form:"email" json:"email" uri:"email" xml:"email" binding:"required"`             // 邮箱
	NickName string `json:"nickname" gorm:"default:''"`                                                // 昵称
}
