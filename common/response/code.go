package response

var (
	OK          = New(0, "success")
	ServerError = New(500, "服务错误")
)
