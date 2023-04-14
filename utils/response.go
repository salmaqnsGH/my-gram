package utils

type ErrorResponse struct {
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
}

func ErrResponse(status int, message string, error interface{}) ErrorResponse {
	response := ErrorResponse{
		Error:   error,
		Status:  status,
		Message: message,
	}

	return response
}
