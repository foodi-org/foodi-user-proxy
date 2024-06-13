package xerror

var MessageType map[int]string

func init() {
	MessageType = map[int]string{}
	MessageType[SuccessCode] = "success"
	MessageType[CommonFailCode] = "common fail"
}
