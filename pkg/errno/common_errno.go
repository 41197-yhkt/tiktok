package errno

var (
	Success                   = NewErrNo(0, "成功")
	ServerError               = NewErrNo(1000, "服务内部错误")
	InvalidParams             = NewErrNo(1001, "入参错误")
	NotFound                  = NewErrNo(1002, "找不到")
	UnauthorizedAuthNotExist  = NewErrNo(1003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewErrNo(1004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewErrNo(1005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewErrNo(1006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewErrNo(1007, "请求过多")
)