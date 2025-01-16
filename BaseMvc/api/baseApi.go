package api

import (
	"github.com/Ideal-f/iCommon/BaseMvc"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
	*BaseMvc.Base
}

func (t *BaseApi) Init(ginContent *gin.Context) {
	// 初始化
	t.InitBase(ginContent)
}
