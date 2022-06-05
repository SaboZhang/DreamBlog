// Package middleware
// @projectName DreamBlog
// @description Google 人机验证中间件,可配置某个方法是否需要Google人机验证
// @author 张涛
// @version 1.0.0
// @createTime 2022年05月22日 15:11:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package middleware

import (
	"dream-blog/global"
	req "dream-blog/model/common/request"
	"dream-blog/model/common/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// 人机验证请求路径
const reCAPTCHAPath = "/recaptcha/api/siteverify"

type netClient interface {
	PostForm(url string, formValues url.Values) (resp *http.Response, err error)
}

type clock interface {
	Since(t time.Time) time.Duration
}

type realClock struct {
}

func (realClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

type ReCAPTCHA struct {
	client        netClient
	Secret        string
	ReCAPTCHALink string
	Timeout       time.Duration
	horloge       clock
}

// RecaptchaVerify
// @Description: Google 人机验证结果校验
//
func RecaptchaVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 是否启用Google人机验证
		if !global.SYS_CONFIG.RecaptchaSettings.Enabled || global.SYS_CONFIG.System.Env == "develop" {
			return
		}
		// 获取参数
		googleRecaptchaToken := ctx.Request.Header.Get(global.SYS_CONFIG.RecaptchaSettings.HeaderKey)
		if googleRecaptchaToken == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "验证参数不存在，人机验证失败！", ctx)
			ctx.Abort()
			return
		}
		r, err := NewReCAPTCHA(global.SYS_CONFIG.RecaptchaSettings.SiteSecret, time.Second*10)
		// 去Google进行验证
		res, err := r.verify(req.ReCHAPTCHARequest{Response: googleRecaptchaToken,
			RemoteIP: ctx.ClientIP(),
			Secret:   global.SYS_CONFIG.RecaptchaSettings.SiteSecret})
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "Google人机验证网络请求失败", ctx)
			ctx.Abort()
			return
		}
		if !res.Success || res.Score != 0 && res.Score < global.SYS_CONFIG.RecaptchaSettings.MinimumScore {
			response.FailWithDetailed(gin.H{"reload": true}, "人机验证失败！", ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// NewReCAPTCHA
// @Description: 人机验证辅助调用方法，生成ReCAPTCHA
//
func NewReCAPTCHA(ReCAPTCHASecret string, timeout time.Duration) (ReCAPTCHA, error) {
	if ReCAPTCHASecret == "" {
		return ReCAPTCHA{}, fmt.Errorf("recaptcha secret cannot be blank")
	}
	return ReCAPTCHA{
		client: &http.Client{
			Timeout: timeout,
		},
		horloge:       &realClock{},
		Secret:        ReCAPTCHASecret,
		ReCAPTCHALink: reCAPTCHAPath,
		Timeout:       timeout,
	}, nil
}

// verify
// @Description: Google 人机验证请求
//
func (r *ReCAPTCHA) verify(req req.ReCHAPTCHARequest) (res response.ReCHAPTCHAResponse, err error) {
	var fromValues url.Values
	if req.RemoteIP == "" || !IsPublicIP(req.RemoteIP) {
		fromValues = url.Values{"secret": {req.Secret}, "response": {req.Response}}
	} else {
		fromValues = url.Values{"secret": {req.Secret}, "response": {req.Response}, "remoteip": {req.RemoteIP}}
	}
	verifyUrl := global.SYS_CONFIG.RecaptchaSettings.VerifyBaseUrl + r.ReCAPTCHALink
	resp, err := r.client.PostForm(verifyUrl, fromValues)
	if err != nil {
		global.SYS_LOG.Error("Google人机验证网络请求失败！", zap.Error(err))
		return res, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			global.SYS_LOG.Error("ReadCloser Exception", zap.Error(err))
		}
	}(resp.Body)
	resultBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.SYS_LOG.Error("Body数据读取异常！", zap.Error(err))
		return res, err
	}
	err = json.Unmarshal(resultBody, &res)
	if err != nil {
		global.SYS_LOG.Error("json数据格式转换异常", zap.Error(err))
		return res, err
	}
	return res, err
}

// IsPublicIP
// @Description: 判断是否公网IP
// @return bool
//
func IsPublicIP(ip string) bool {
	n, _ := exnet.IPString2Long(ip)
	i, _ := exnet.Long2IP(n)
	if i.IsLoopback() || i.IsLinkLocalMulticast() || i.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := i.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
