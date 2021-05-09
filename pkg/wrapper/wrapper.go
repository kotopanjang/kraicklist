package wrapper

// Result is a struct to return from function
type Result struct {
	IsError bool
	Err     error
	Data    interface{}
}

// httpResponse is a struct to return to HTTP Request
type httpResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}
