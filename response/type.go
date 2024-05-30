package response

type (
	Body struct {
		Code int `json:"code"`

		Msg string `json:"msg"`

		Data interface{} `json:"data"`
	}
)
