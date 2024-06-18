package response

var (

	// SuccessResponse 统一成功返回
	SuccessResponse = New(0, "success")

	// ServerError 服务错误统一返回
	ServerError = New(500, "服务错误")

	BindParamError = New(1100010, "param bind error")
)
