package common

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Uri struct {
	Method   string // 请求方法
	Path     string // 请求url
	Project  string // 项目名称
	Version  string // 接口版本号
	TenantId uint   // 请求的租户ID
}

// RenderUri 解析Uri信息
func (t *Uri) RenderUri(ginContext *gin.Context) {
	// 初始化uri信息

	t.Path = ginContext.Request.RequestURI

	// 获取项目名称
	uriArr := strings.Split(t.Path, "/")
	if len(uriArr) < 4 {
		t.Project = uriArr[0]
		t.Version = uriArr[1]
	} else {
		t.Project = t.Path
	}

	return
}
