package my

import (
	"github.com/Ideal-f/iCommon/BaseController"
	"github.com/gin-gonic/gin"
)

type BaseMy struct {
	*BaseController.Base
}

func (t *BaseMy) Init(ginContent *gin.Context) {
	// 初始化
	t.InitBase(ginContent)
}
