package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type Token struct {
	token          string // 当前使用token
	oldToken       string // 旧token (刷新之前的token)
	CreateTime     int64  // 创建时间
	CreateTimeStr  string // 创建时间字符串
	ExpireTime     int64  // 过期时间
	ExpireTimeStr  string // 过期时间字符串
	LastActiveTime int64  // 最后活跃时间
}

const (
	tokenExpireSpan = 28800 // token过期时间 单位:秒 默认8小时

	tokenRedisCacheSpan = 7200 // redis缓存时间 单位:秒 默认2小时

	tokenRefreshSpan = 3600 // 过期多久可以进行刷新 单位:秒 默认1小时
)

// RenderToken 解析token
func (t *Token) RenderToken(GinContent *gin.Context) {

	// 获取header中的token
	t.token = GinContent.GetHeader("token")

	// 从redis里面获取token信息

	// 设置最后活跃时间

	return
}

// RefreshToken 刷新token
func (t *Token) RefreshToken() (err error) {

	// 失去活跃多久，可以刷新token
	if (t.LastActiveTime + tokenRefreshSpan) < time.Now().Unix() {
		err = errors.New("token刷新超时")
		return
	}

	// 根据当前token信息生成新的token

	// 设置旧token的最后过期时间

	return
}

// checkToken 检查token有效期
func (t *Token) checkToken(isAutoRefresh bool) (errCode int, err error) {
	// 检查token是否为空
	if t.token == "" {
		errCode = 100401
		err = errors.New("请先登录，登录信息为空！")
		return
	}

	// 检查token是否过期
	if t.ExpireTime < time.Now().Unix() {
		if isAutoRefresh {
			// 刷新token
			err = t.RefreshToken()
			if err != nil {
				errCode = 100401
				return
			}
		} else {
			errCode = 100410
			err = errors.New("token已过期")
			return
		}
	}

	return
}
