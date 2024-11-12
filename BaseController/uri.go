package BaseController

import "strings"

type Uri struct {
	FullUri string `json:"full_uri"`
	Project string `json:"project"`
}

// RenderUri 解析Uri信息
func (t *Base) RenderUri() {
	// 初始化header信息
	t.Uri = &Uri{}

	t.Uri.FullUri = t.GinContent.Request.RequestURI

	// 获取项目名称
	uriArr := strings.Split(t.Uri.FullUri, "/")
	if len(uriArr) > 1 {
		t.Uri.Project = uriArr[0]
	} else {
		t.Uri.Project = t.Uri.FullUri
	}

	return
}
