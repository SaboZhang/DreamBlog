// Package system
// @projectName DreamBlog
// @description 黑名单结构体，实现登录Token的过期
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 13:06:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	"dream-blog/model"
	"time"
)

type BlackRecord struct {
	model.BaseUUID
	Jti          string    `gorm:"type:text;comment:jwt"`
	UserName     string    `gorm:"type:varchar(50);comment:用户名"`
	CreateUserId int64     `gorm:"comment:创建人"`
	CreateTime   time.Time `gorm:"comment:创建时间"`
}

// TableName
// @Description: 指定数据库表名
// @receiver BlackRecord
// @return string 数据库表名
//
func (BlackRecord) TableName() string {
	return "black_record"
}
