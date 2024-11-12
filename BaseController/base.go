package BaseController

import "github.com/gin-gonic/gin"

type Base struct {
	GinContent *gin.Context
	Uri        *Uri
	Header     *Header
	Token      *Token
}

func (t *Base) InitBase(ginContent *gin.Context) {
	// 赋值上下文信息
	t.GinContent = ginContent

	// 解析token信息
	t.RenderToken()

	// 解析uri信息
	t.RenderHeader()

	// 解析header信息
	t.RenderHeader()

}
