package api

import (
	"github.com/gin-gonic/gin"
	"iCommon/BaseController"
)

type BaseApi struct {
	*BaseController.Base
}

func (t *BaseApi) Init(ginContent *gin.Context) {
	// 初始化
	t.InitBase(ginContent)
}
