package token

type TokenError struct {
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误描述
}

func (e *TokenError) Error() string {
	return e.Message
}

func NewTokenError(code, message string) *TokenError {
	return &TokenError{
		Code:    code,
		Message: message,
	}
}
