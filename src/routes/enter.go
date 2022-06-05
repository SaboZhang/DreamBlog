// Package routes
/**
 * @projectName DreamBlog
 * @author 张涛
 * @version 1.0.0
 * @description 全局路由配置
 * @createTime 2022年05月02日 20:43:00
 * @lastEditors: 张涛
 * @lastEditTime:
 * 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
 */
package routes

import (
	"dream-blog/routes/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
