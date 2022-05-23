// Package config
/**
 * @projectName DreamBlog
 * @author 张涛
 * @version 1.0.0
 * @description 数据库配置
 * @createTime 2022年05月02日 23:40:00
 * @lastEditors: 张涛
 * @lastEditTime:
 * 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
 */
package config

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                         // 服务器地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                         // 端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                   // 高级配置
	Dbname       string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`                   // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`             // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`             // 数据库密码
	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"logMode" json:"logMode" yaml:"logMode"`                // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"logZap" json:"logZap" yaml:"logZap"`                   // 是否通过zap写入日志文件
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
