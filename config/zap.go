// Package config
// @projectName DreamBlog
// @description zap 日志
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月04日 22:27:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package config

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                         // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                      // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                      // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`               // 日志文件夹
	ShowLine      bool   `mapstructure:"showLine" json:"showLine" yaml:"showLine"`                // 显示行
	EncodeLevel   string `mapstructure:"encodeLevel" json:"encodeLevel" yaml:"encodeLevel"`       // 编码级
	StacktraceKey string `mapstructure:"stacktraceKey" json:"stacktraceKey" yaml:"stacktraceKey"` // 栈名
	LogInConsole  bool   `mapstructure:"logInConsole" json:"logInConsole" yaml:"logInConsole"`    // 输出控制台
}
