package httperror

import (
	stderrors "errors"
	"net/http"

	"github.com/Thinura/go-eventops-platform/internal/apperror"
)

type Error struct {
	StatusCode int
	Code       string
	Message    string
}

func Map(err error) Error {
	var appErr *apperror.Error

	if stderrors.As(err, &appErr) {
		return mapAppError(appErr)
	}

	return Error{
		StatusCode: http.StatusInternalServerError,
		Code:       string(apperror.CodeInternal),
		Message:    "internal server error",
	}
}

func mapAppError(err *apperror.Error) Error {
	switch err.Code {
	case apperror.CodeInvalidRequestBody,
		apperror.CodeValidation,
		apperror.CodeEventSourceRequired,
		apperror.CodeEventTypeRequired,
		apperror.CodeEventEntityIDRequired,
		apperror.CodeEventOccurredAtMissing,
		apperror.CodeUnsupportedEventType:
		return Error{
			StatusCode: http.StatusBadRequest,
			Code:       string(err.Code),
			Message:    err.Message,
		}

	case apperror.CodeNotFound:
		return Error{
			StatusCode: http.StatusNotFound,
			Code:       string(err.Code),
			Message:    err.Message,
		}

	case apperror.CodeConflict,
		apperror.CodeEventAlreadyExists:
		return Error{
			StatusCode: http.StatusConflict,
			Code:       string(err.Code),
			Message:    err.Message,
		}

	default:
		return Error{
			StatusCode: http.StatusInternalServerError,
			Code:       string(apperror.CodeInternal),
			Message:    "internal server error",
		}
	}
}
