// Package utils
// @projectName DreamBlog
// @description 文件目录工具类
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月04日 22:30:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package utils

import (
	"errors"
	"os"
)

// PathExists
// @Description: 判断文件是否存在
// @return bool
// @return error
//
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
