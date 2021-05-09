package wrapper

// Result is a struct that contains property of result of function's return
type Result struct {
	IsError bool
	Err     error
	Data    interface{}
}

type httpResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}
