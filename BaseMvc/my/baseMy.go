package my

import (
	"github.com/Ideal-f/iCommon/BaseMvc"
	"github.com/gin-gonic/gin"
)

type BaseMy struct {
	*BaseMvc.Base
}

func (t *BaseMy) Init(ginContent *gin.Context) {
	// 初始化
	t.InitBase(ginContent)
}
