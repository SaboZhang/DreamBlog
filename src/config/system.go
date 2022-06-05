// Package config
// @projectName DreamBlog
// @description 系统环境结构体
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月05日 01:37:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                               // 环境值
	Addr          int    `mapstructure:"port" json:"port" yaml:"port"`                            // 端口值
	DbType        string `mapstructure:"dbType" json:"dbType" yaml:"dbType"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"ossType" json:"ossType" yaml:"ossType"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"useMultipoint" json:"useMultipoint" yaml:"useMultipoint"` // 多点登录拦截
	UseRedis      bool   `mapstructure:"useRedis" json:"useRedis" yaml:"useRedis"`                // 使用redis
	LimitCountIP  int    `mapstructure:"iplimitCount" json:"iplimitCount" yaml:"iplimitCount"`
	LimitTimeIP   int    `mapstructure:"iplimitTime" json:"iplimitTime" yaml:"iplimitTime"`
	UseTls        bool   `mapstructure:"useTls" json:"useTls" yaml:"useTls"`
	CertFile      string `mapstructure:"certFile" json:"certFile" yaml:"certFile"`
	KeyFile       string `mapstructure:"keyFile" json:"keyFile" yaml:"keyFile"`
}
