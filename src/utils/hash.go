// Package utils
// @projectName DreamBlog
// @description 密码加密验证工具
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 17:41:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package utils

import "golang.org/x/crypto/bcrypt"

// BcryptHash
// @Description: 使用 bcrypt 对密码进行加密
// @return string
//
func BcryptHash(password string, salt string) string {
	saltPassword := password + salt
	bytes, _ := bcrypt.GenerateFromPassword([]byte(saltPassword), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck
// @Description: 对比明文密码和数据库的哈希值
// @return bool
//
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
