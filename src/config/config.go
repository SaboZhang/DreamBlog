// Package config
// @projectName DreamBlog
// @description 服务器配置
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月04日 20:48:00
// @lastEditors: 张涛
// @lastEditTime: 2022年5月24日00:20:00
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

type Server struct {
	JWT               JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`                                           //JWT
	Mysql             Mysql     `mapstructure:"connection" json:"connection" yaml:"connection"`                      // gorm
	Zap               Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`                                           // zap日志
	System            System    `mapstructure:"server" json:"server" yaml:"server"`                                  // 系统配置
	Casbin            Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`                                  //casbin配置
	Local             Local     `mapstructure:"storage" json:"storage" yaml:"storage"`                               // 本地存储配置
	Qiniu             Qiniu     `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`                                     // 七牛云配置
	Cors              CORS      `mapstructure:"cors" json:"cors" yaml:"cors"`                                        // 跨域配置
	RecaptchaSettings Recaptcha `mapstructure:"recaptchaSettings" json:"recaptchaSettings" yaml:"recaptchaSettings"` // Google人机验证
	Captcha           Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`                               // 普通图片验证码配置
	Redis             Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`                                     // Redis
}
