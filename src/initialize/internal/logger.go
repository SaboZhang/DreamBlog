// Package internal
// @projectName DreamBlog
// @description TODO
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月04日 22:00:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package internal

import (
	"dream-blog/global"
	"fmt"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter
// @Description: 构造函数
// @return *writer
//
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf
// @Description: 格式化打印日志
// @receiver w
//
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	logZap = global.SYS_CONFIG.Mysql.LogZap
	if logZap {
		global.SYS_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
