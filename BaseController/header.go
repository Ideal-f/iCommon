package BaseController

type Header struct {
	StaffId     int          `json:"staff_id"`
	StaffName   string       `json:"staff_name"`
	CompanyId   int          `json:"company_id"`
	CompanyName string       `json:"company_name"`
	Departments []Department `json:"departments"`
	Roles       []Role       `json:"roles"`
}
type Department struct {
	DepartmentId   int    `json:"department_id"`
	DepartmentName string `json:"department_name"`
}
type Role struct {
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

const (
	headerRedisCacheSpan = 43200 // redis缓存时间 单位:秒 默认12小时
)

// RenderHeader 解析header
func (t *Base) RenderHeader() {
	// 初始化header信息
	t.Header = &Header{}

	// 如果没有token信息则返回
	if t.Token.token == "" {
		return
	}

	// 根据token信息获取redis用户信息

	return
}
