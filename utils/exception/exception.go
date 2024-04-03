package exception

type Exception struct {
	message    string
	statusCode int
}

func New(message string, statusCode int) *Exception {
	e := Exception{
		message:    message,
		statusCode: statusCode,
	}

	return &e
}
