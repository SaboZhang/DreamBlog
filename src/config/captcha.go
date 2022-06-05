// Package config
// @projectName DreamBlog
// @description 验证码配置
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月22日 21:28:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

type Captcha struct {
	KeyLong   int  `mapstructure:"keyLong" json:"keyLong" yaml:"keyLong"`       // 验证码长度
	ImgWidth  int  `mapstructure:"imgWidth" json:"imgWidth" yaml:"imgWidth"`    // 验证码宽度
	ImgHeight int  `mapstructure:"imgHeight" json:"imgHeight" yaml:"imgHeight"` // 验证码高度
	Enable    bool `mapstructure:"enable" json:"enable" yaml:"enable"`          // 是否启用图片验证码
}
