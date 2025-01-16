package api

import (
	"github.com/Ideal-f/iCommon/BaseController"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
	*BaseController.Base
}

func (t *BaseApi) Init(ginContent *gin.Context) {
	// 初始化
	t.InitBase(ginContent)
}
