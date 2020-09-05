package common

// Header Key
const (
	HeaderContentType  = "Content-Type"
	HeaderOrigin       = "Origin"
	HeaderReferer      = "Referer"
	HeaderSecFetchMode = "Sec-Fetch-Mode"
	HeaderXCSRFTOKEN   = "X-CSRF-TOKEN"
)

type ContentType string

func (c ContentType) String() string {
	return string(c)
}

// Content-Type HTTP Request
const (
	ContentTypeBinary ContentType = "application/octet-stream"
	ContentTypeForm   ContentType = "application/x-www-form-urlencoded"
	ContentTypeJSON   ContentType = "application/json"
	ContentTypeHTML   ContentType = "text/html; charset=utf-8"
	ContentTypeText   ContentType = "text/plain; charset=utf-8"
)

// Url Constants
const (
	OriginAPISMMO = "https://api.simple-mmo.com"
	RefererTravel = OriginAPISMMO + "/travel"
)

// Messages
const (
	FailedToMarshal       = "failed to marshal"
	FailedToWriteResponse = "failed to write response"
)

// Param Key
const (
	ParamToken    = "token"
	ParamAPIToken = "api_token"
	ParamHash     = "hash"
)
