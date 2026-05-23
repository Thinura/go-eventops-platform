package apperror

type Code string

const (
	CodeValidation Code = "VALIDATION_ERROR"
	CodeNotFound   Code = "NOT_FOUND"
	CodeConflict   Code = "CONFLICT"
	CodeInternal   Code = "INTERNAL_ERROR"

	CodeInvalidRequestBody Code = "INVALID_REQUEST_BODY"

	CodeEventSourceRequired    Code = "EVENT_SOURCE_REQUIRED"
	CodeEventTypeRequired      Code = "EVENT_TYPE_REQUIRED"
	CodeEventEntityIDRequired  Code = "EVENT_ENTITY_ID_REQUIRED"
	CodeEventOccurredAtMissing Code = "EVENT_OCCURRED_AT_REQUIRED"
	CodeUnsupportedEventType   Code = "UNSUPPORTED_EVENT_TYPE"

	CodeEventPublishFailed Code = "EVENT_PUBLISH_FAILED"
	CodeEventSaveFailed    Code = "EVENT_SAVE_FAILED"
	CodeEventAlreadyExists Code = "EVENT_ALREADY_EXISTS"
)
