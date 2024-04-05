package response

type Response struct {
	Message string `json:"message"`
}

func New(message string) *Response {
	return &Response{
		Message: message,
	}
}
