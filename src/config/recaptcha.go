// Package config
// @projectName DreamBlog
// @description Google 人机验证配置
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月22日 15:12:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

type Recaptcha struct {
	HeaderKey     string  `mapstructure:"headerKey" json:"headerKey" yaml:"headerKey"`
	Enabled       bool    `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	MinimumScore  float32 `mapstructure:"minimumScore" json:"minimumScore" yaml:"minimumScore"`
	SiteKey       string  `mapstructure:"siteKey" json:"siteKey" yaml:"siteKey"`
	SiteSecret    string  `mapstructure:"siteSecret" json:"siteSecret" yaml:"siteSecret"`
	VerifyBaseUrl string  `mapstructure:"verifyBaseUrl" json:"verifyBaseUrl" yaml:"verifyBaseUrl"`
}
