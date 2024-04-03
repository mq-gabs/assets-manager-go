package exception

type Exception struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func New(message string, statusCode int) *Exception {
	e := Exception{
		Message:    message,
		StatusCode: statusCode,
	}

	return &e
}
