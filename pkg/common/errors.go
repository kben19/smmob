package common

import "errors"

// Parse and cast errors
var (
	ErrInvalidCastType    = errors.New("invalid cast type")
	ErrInvalidPayload     = errors.New("invalid payload value")
	ErrInvalidHeaderValue = errors.New("invalid header value")
)

// HTTP request errors
var (
	ErrInvalidContentTypeHeader = errors.New("invalid content type header")
)
