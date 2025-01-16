package BaseMvc

import (
	"github.com/Ideal-f/iCommon/BaseMvc/common"
	"github.com/gin-gonic/gin"
)

type Base struct {
	GinContent *gin.Context
	Uri        *common.Uri
	Header     *common.Header
	Token      *common.Token
}

func (t *Base) InitBase(ginContent *gin.Context) {
	// 赋值上下文信息
	t.GinContent = ginContent

	// 解析token信息
	if t.Token == nil {
		t.Token = new(common.Token)
	}
	t.Token.RenderToken(t.GinContent)

	// 解析uri信息
	if t.Uri == nil {
		t.Uri = new(common.Uri)
	}
	t.Uri.RenderUri(t.GinContent)

	// 解析header信息
	if t.Header == nil {
		t.Header = new(common.Header)
	}
	t.Header.RenderHeader(t.GinContent)

}
