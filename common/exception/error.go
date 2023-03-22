package exception

// ApiError interface unify error
type ApiError struct {
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误描述
}

func (e *ApiError) Error() string {
	return e.Message
}

// NewApiError create ApiError
func NewApiError(code, message string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
	}
}
