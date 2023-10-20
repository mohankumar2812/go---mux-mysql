package utils

type ErrorMessages struct {
	Message string `json message`
	Status  int    `json status`
	Error   string `json error`
}

func BadRequest(message string) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Status: 500,
		Error: "Bad Request",
	}
}

func NotFound(message string) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Status: 404,
		Error: "Not Found",
	}
}

func InternalErr(message string) *ErrorMessages {
	return &ErrorMessages{
		Message: message,
		Status: 400,
		Error: "Internal Server Error",
	}
}
