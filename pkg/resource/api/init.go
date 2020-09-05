package api

import (
	"io"
	"net/http"
)

const (
	logTagResourceAPI = " [ResourceAPI]"
)

type ResourceAPI struct {
	httpClient *http.Client
}

type ResourceAPIItf interface {
	HTTPRequest(method string, url string, body io.Reader, urlQuery map[string]string, header map[string]string) ([]byte, error)
	GetHTTPRequest(url string, queries map[string]string, header map[string]string) ([]byte, error)
	PostHTTPRequest(url string, queries map[string]string, header map[string]string, bodyRaw interface{}) ([]byte, error)
}

func InitResourceAPI(httpClient *http.Client) ResourceAPIItf {
	return &ResourceAPI{httpClient: httpClient}
}
