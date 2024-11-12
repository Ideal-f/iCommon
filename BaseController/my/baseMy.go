package my

import (
	"github.com/gin-gonic/gin"
	"iCommon/BaseController"
)

type BaseMy struct {
	*BaseController.Base
}

func (t *BaseMy) Init(ginContent *gin.Context) {
	// 初始化
	t.InitBase(ginContent)
}
