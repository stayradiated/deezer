package deezer

import "fmt"

const (
	ERR_QUOTA             = 4
	ERR_PERMISSION        = 200
	ERR_TOKEN_INVALID     = 300
	ERR_PARAMETER         = 500
	ERR_PARAMETER_MISSING = 501
	ERR_QUERY_INVALID     = 600
	ERR_SERVICE_BUSY      = 700
	ERR_DATA_NOT_FOUND    = 800
)

type Error struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type ErrorResponse struct {
	Error Error `json:"error,omitempty"`
}

func (e Error) String() string {
	return fmt.Sprintf(
		"%d: %s - %s",
		e.Code, e.Type, e.Message,
	)
}
