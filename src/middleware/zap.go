// Package middleware
// @projectName DreamBlog
// @description TODO
// @author 张涛
// @version 1.0.0
// @createTime 2022年08月22日 22:31:00
// @lastEditors: 张涛
// @lastEditTime:
// 世界上最遥远的距离不是生与死，而是你亲手制造的BUG就在你眼前，你却怎么都找不到她
// @Copyright (c) 2022 by Rick email: tao993859833@live.cn, All Rights Reserved
//
package middleware

import (
	"bytes"
	"dream-blog/global"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// GinLogger
// @Description: GIN框架日志的默认接收器
// @return gin.HandlerFunc
//
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		post := ""
		if c.Request.Method == "POST" {
			// 把request的内容读取出来
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			err := c.Request.Body.Close()
			if err != nil {
				return
			}
			// 把刚刚读出来的再写进去
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			switch c.ContentType() {
			case "application/json":
				var result map[string]interface{}
				d := jsoniter.NewDecoder(bytes.NewReader(bodyBytes))
				d.UseNumber()
				if err := d.Decode(&result); err == nil {
					bt, _ := jsoniter.Marshal(result)
					post = string(bt)
				}
			default:
				post = string(bodyBytes)
			}
		}
		c.Next()
		cost := time.Since(start)
		global.SYS_LOG.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.Duration("cost", cost),
			zap.String("method", c.Request.Method),
			zap.String("api", path),
			zap.String("query", query),
			zap.String("post", post),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}

}

func GinRecoveryLogger(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.SYS_LOG.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					global.SYS_LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					global.SYS_LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
