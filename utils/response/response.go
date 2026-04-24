package common

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  []string   `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Errors:  []string{err},
		Data:    data,
	}
}