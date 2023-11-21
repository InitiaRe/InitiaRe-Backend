package constant

const (
	// ERROR
	STATUS_CODE_BAD_REQUEST     = "BAD_REQUEST"
	STATUS_CODE_NOT_FOUND       = "NOT_FOUND"
	STATUS_CODE_UNAUTHORIZED    = "UNAUTHORIZED"
	STATUS_CODE_FORBIDDEN       = "FORBIDDEN"
	STATUS_CODE_INTERNAL_SERVER = "INTERNAL_SERVER_ERROR"
	STATUS_CODE_REQUEST_TIMEOUT = "REQUEST_TIMEOUT"
)

const (
	// SUCCESS
	STATUS_MESSAGE_OK       = "OK"
	STATUS_MESSAGE_CREATED  = "CREATED"
	STATUS_MESSAGE_ACCEPTED = "ACCEPTED"

	// ERROR
	STATUS_MESSAGE_EMAIL_ALREADY_EXISTS      = "EMAIL_ALREADY_EXISTS"
	STATUS_MESSAGE_INVALID_GENDER_TYPE       = "INVALID_GENDER_TYPE"
	STATUS_MESSAGE_USER_NOT_FOUND            = "USER_NOT_FOUND"
	STATUS_MESSAGE_INVALID_EMAIL_OR_PASSWORD = "INVALID_EMAIL_OR_PASSWORD"
	STATUS_MESSAGE_INTERNAL_SERVER_ERROR     = "AN ERROR OCCURRED WHILE PROCESSING YOUR REQUEST"
	STATUS_MESSAGE_INVALID_JWT_TOKEN         = "INVALID_JWT_TOKEN"
	STATUS_MESSAGE_USER_INACTIVE             = "USER_INACTIVE"
)
