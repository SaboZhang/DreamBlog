// Package config
// @projectName DreamBlog
// @description 七牛云配置
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月21日 22:27:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                            // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                      // 空间名称
	ImgPath       string `mapstructure:"imgPath" json:"imgPath" yaml:"imgPath"`                   // CDN加速域名
	UseHTTPS      bool   `mapstructure:"useHttps" json:"useHttps" yaml:"useHttps"`                // 是否使用https
	AccessKey     string `mapstructure:"accessKey" json:"accessKey" yaml:"accessKey"`             // 秘钥AK
	SecretKey     string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`             // 秘钥SK
	UseCdnDomains bool   `mapstructure:"useCdnDomains" json:"useCdnDomains" yaml:"useCdnDomains"` // 上传是否使用CDN上传加速
}
