// Package core
// @projectName DreamBlog
// @description 系统服务
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月05日 01:42:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package core

import (
	"dream-blog/global"
	"dream-blog/initialize"
	"dream-blog/service/system"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	if global.SYS_CONFIG.System.UseRedis || global.SYS_CONFIG.System.UseMultipoint {
		// 初始化Redis
		initialize.Redis()
	}
	// 从db加载jwt数据
	if global.SYS_DB != nil {
		system.LoadAll()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.SYS_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.SYS_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	欢迎使用DreamBlog
	当前版本v1.0.0
	默认自动化文档地址:http(s)://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8088
`, address)
	// https 模式
	if global.SYS_CONFIG.System.UseTls {
		err := Router.RunTLS(address, global.SYS_CONFIG.System.CertFile, global.SYS_CONFIG.System.KeyFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	global.SYS_LOG.Error(s.ListenAndServe().Error())
}
