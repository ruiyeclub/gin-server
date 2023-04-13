package enum

var (
	PasswdErr      = NewErrorEnum("0001", "密码错误")
	NotUser        = NewErrorEnum("0002", "用户不存在")
	DataErr        = NewErrorEnum("0003", "数据服务异常")
	UserIsExist    = NewErrorEnum("0004", "用户已存在")
	ValidateError  = NewErrorEnum("0005", "非授权用户")
	ArgumentErr    = NewErrorEnum("0006", "参数类型异常")
	LoginErr       = NewErrorEnum("0007", "请授权登录")
	AuthErr        = NewErrorEnum("0008", "认证失败")
	TokenInvalid   = NewErrorEnum("0010", "无效token")
	RedisErr       = NewErrorEnum("0011", "Redis操作失败")
	JsonErr        = NewErrorEnum("0012", "Json操作失败")
	DataExisted    = NewErrorEnum("0013", "数据已存在")
	DataNotExisted = NewErrorEnum("0014", "数据不存在")
	PermissionsErr = NewErrorEnum("0015", "权限不足")
	TokenExpires   = NewErrorEnum("0016", "token过期")
)

func NewErrorEnum(code string, msg string) *ErrorEnum {
	return &ErrorEnum{code, msg}
}

type ErrorEnum struct {
	Code string
	Msg  string
}
