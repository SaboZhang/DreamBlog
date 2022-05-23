// Package config
// @projectName DreamBlog
// @description 本地文件路径配置
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月21日 22:26:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

type Local struct {
	Path string `mapstructure:"prefixPath" json:"prefixPath" yaml:"prefixPath"` // 本地文件路径
}
