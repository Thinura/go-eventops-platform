package apperror

func New(code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func Wrap(code Code, message string, cause error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

func Validation(code Code, message string) *Error {
	return New(code, message)
}

func NotFound(message string) *Error {
	return New(CodeNotFound, message)
}

func Conflict(code Code, message string) *Error {
	return New(code, message)
}

func Internal(message string, cause error) *Error {
	return Wrap(CodeInternal, message, cause)
}
