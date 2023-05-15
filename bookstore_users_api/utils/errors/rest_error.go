package restErrors

import "net/http"

type RestErr struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

func UserNotFoundError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not found",
	}
}

func UserAlreadyExistsError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusConflict,
		Error:      "conflict",
	}

}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}

}
