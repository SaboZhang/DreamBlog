/**
 * @projectName DreamBlog
 * @author 张涛
 * @version 1.0.0
 * @description 程序主入口
 * @createTime 2022年05月02日 18:23:00
 * @lastEditors: 张涛
 * @lastEditTime:
 * 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
 */
package main

import (
	"database/sql"
	"dream-blog/core"
	"dream-blog/global"
	"dream-blog/initialize"
	"go.uber.org/zap"
)

// @title DreamBlog API Golang
// @version v1
// @description DreamBlog Golang 版本swagger
// @securityDefinitions.apikey ApiKeyAuth
// @JWT授权(数据将在请求头中进行传输) 参数结构: "Authorization: {token}"
// @name Authorization
// @in header
// @BasePath /
//
func main() {
	global.SYS_VP = core.Viper() // 初始化Viper
	global.SYS_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.SYS_LOG)
	global.SYS_DB = initialize.GormMysql() // gorm连接数据库
	if global.SYS_DB != nil {
		initialize.RegisterTables(global.SYS_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.SYS_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.SYS_LOG.Error("数据库异常！！!", zap.Error(err))
				return
			}
		}(db)
	}
	core.RunWindowsServer()
}
