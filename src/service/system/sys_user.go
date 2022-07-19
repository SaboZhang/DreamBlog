// Package system
// @projectName DreamBlog
// @description 用户管理
// @author 张涛
// @version 1.0.0
// @createTime 2022年06月15日 11:27:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	"dream-blog/global"
	"dream-blog/model/system"
	"go.uber.org/zap"
)

// GetInformation
// @Description: 获取单个用户或者当前登陆用户的信息
// @Receiver userService
// @Router
//
func (userService *UserService) GetInformation(uid uint64) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.SYS_DB.First(&reqUser, "id = ?", uid).Error
	if err != nil {
		global.SYS_LOG.Error("用户不存在", zap.Error(err))
		return err, reqUser
	}
	return nil, reqUser
}
