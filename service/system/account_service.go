// Package system
// @projectName DreamBlog
// @description 登陆|注册逻辑处理
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月15日 09:45:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package system

import (
	"dream-blog/global"
	"dream-blog/model/system"
	systemReq "dream-blog/model/system/request"
	systemResp "dream-blog/model/system/response"
	"dream-blog/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type UserService struct {
}

// Login
// @Description: 用户登陆service
// @Receiver userService
// @return err
// @return userInter
//
func (userService *UserService) Login(u systemReq.Login) (err error, userInter *systemResp.RespUser) {
	if nil == global.SYS_DB {
		return fmt.Errorf("db not init"), nil
	}
	var user system.SysUser
	var ident system.SysUserIdentity
	err = global.SYS_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		saltPassword := u.Password + user.Salt
		global.SYS_DB.Where("create_user_id = ?", user.ID).First(&ident)
		if ok := utils.BcryptCheck(saltPassword, ident.Credential); !ok {
			return errors.New("密码错误"), nil
		}
	}
	refers := utils.AddRefreshToken()
	user.RefreshToken = refers
	user.LastLoginTime = time.Now()
	global.SYS_DB.Save(&user)
	var res systemResp.RespUser
	res.ID = user.ID
	res.Nickname = user.Nickname
	res.UUID = ident.ID
	res.Email = user.Email
	res.Username = user.Username
	res.RefreshToken = refers
	return err, &res

}

// Register
// @Description: 用户注册service
// @Receiver userService
// @return err
// @return userInter
//
func (userService *UserService) Register(u systemReq.Register) (err error, userInter system.SysUser) {
	var user system.SysUser
	var ident system.SysUserIdentity
	if errors.Is(global.SYS_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 创建用户
	user.Username = u.Username
	salt := uuid.New().String()
	user.Salt = salt
	user.Nickname = u.NickName
	user.Email = u.Email
	ident.Credential = utils.BcryptHash(u.Password, salt)
	ident.IdentityType = "Password"
	ident.Identifier = u.Username
	err = global.SYS_DB.Create(&user).Error
	if err == nil {
		ident.Base.CreateUserId = user.ID
		err = global.SYS_DB.Create(&ident).Error
	}
	return err, user
}

func (userService *UserService) ChangePassword(u *systemReq.ChangePasswordStruct) (err error) {
	var user system.SysUser
	var ident system.SysUserIdentity
	err = global.SYS_DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return err
	}
	err = global.SYS_DB.Where("create_user_id = ?", user.ID).First(&ident).Error
	if err != nil {
		return err
	}
	saltPassword := u.Password + user.Salt
	if ok := utils.BcryptCheck(saltPassword, ident.Credential); !ok {
		return errors.New("原密码错误")
	}
	ident.Credential = utils.BcryptHash(u.NewPassword, user.Salt)
	err = global.SYS_DB.Save(ident).Error
	return err
}
