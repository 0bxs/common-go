package status

// 系统统一返回code
const (
	OK               = int8(1) // 请求成功
	SystemError      = int8(2) // 服务器繁忙,请稍后再试
	TokenValid       = int8(3) // Token 过期
	RefreshValid     = int8(4) // 账户已在其他设备登录
	TooEarly         = int8(5) // 刷新太早
	NotLogin         = int8(6) // 未登录
	PermissionDenied = int8(7) // 没有权限
	CustomCode       = int8(8) // 自定义异常
	ParamCode        = int8(9) // 请求参数异常
)
