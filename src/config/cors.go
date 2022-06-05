// Package config
// @projectName DreamBlog
// @description 跨域放行规则配置
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月21日 22:45:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

type CORS struct {
	Mode      string          `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allowOrigin" json:"allowOrigin" yaml:"allowOrigin"`
	AllowMethods     string `mapstructure:"allowMethods" json:"allowMethods" yaml:"allowMethods"`
	AllowHeaders     string `mapstructure:"allowHeaders" json:"allowHeaders" yaml:"allowHeaders"`
	ExposeHeaders    string `mapstructure:"exposeHeaders" json:"exposeHeaders" yaml:"exposeHeaders"`
	AllowCredentials bool   `mapstructure:"allowCredentials" json:"allowCredentials" yaml:"allowCredentials"`
}
