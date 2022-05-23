// Package config
/**
 * @projectName DreamBlog
 * @author 张涛
 * @version 1.0.0
 * @description casbin 配置文件
 * @createTime 2022年05月03日 00:10:00
 * @lastEditors: 张涛
 * @lastEditTime:
 * 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
 */
package config

type Casbin struct {
	ModelPath string `mapstructure:"modelPath" json:"modelPath" yaml:"modelPath"` // 存放casbin模型的相对路径
}
