package api

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/kben19/smmob/pkg/common"
)

var (
	ErrInvalidBody = errors.New("invalid body payload")
)

const (
	logTagHTTPRequest     = logTagResourceAPI + "[HTTPRequest]"
	logTagGetHTTPRequest  = logTagResourceAPI + "[GetHTTPRequest]"
	logTagPostHTTPRequest = logTagResourceAPI + "[PostHTTPRequest]"
)

func (r ResourceAPI) HTTPRequest(
	method string,
	url string,
	body io.Reader,
	urlQuery map[string]string,
	header map[string]string,
) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err.Error() + logTagHTTPRequest)
		return nil, err
	}

	// Url Queries
	query := req.URL.Query()
	for key, value := range urlQuery {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	// Headers
	for key, value := range header {
		req.Header.Add(key, value)
	}

	// Do HTTP Request
	response, err := r.httpClient.Do(req)
	if err != nil {
		log.Println(err.Error() + logTagHTTPRequest)
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error() + logTagHTTPRequest)
		return nil, err
	}
	return bytes, nil
}

func (r ResourceAPI) GetHTTPRequest(url string, queries map[string]string, header map[string]string) ([]byte, error) {
	bytes, err := r.HTTPRequest(http.MethodGet, url, nil, queries, header)
	if err != nil {
		log.Println(err.Error() + logTagGetHTTPRequest)
		return nil, err
	}
	return bytes, nil
}

func (r ResourceAPI) PostHTTPRequest(url string, queries map[string]string, header map[string]string, bodyRaw interface{}) ([]byte, error) {
	body, err := parseBody(bodyRaw, header)
	if err != nil {
		log.Println(err.Error() + logTagPostHTTPRequest)
		return nil, err
	}

	bytes, err := r.HTTPRequest(http.MethodPost, url, body, queries, header)
	if err != nil {
		log.Println(err.Error() + logTagPostHTTPRequest)
		return nil, err
	}
	return bytes, nil
}

func parseBody(body interface{}, header map[string]string) (io.Reader, error) {
	rvalue := reflect.ValueOf(body)
	switch rvalue.Kind() {
	case reflect.Map:
		if rvalue.IsNil() {
			return nil, ErrInvalidBody
		}
		if header[common.HeaderContentType] != common.ContentTypeForm.String() {
			return nil, common.ErrInvalidContentTypeHeader
		}
		urlValues := url.Values{}
		bodyVal, ok := body.(map[string]string)
		if !ok {
			return nil, common.ErrInvalidCastType
		}
		for key, val := range bodyVal {
			urlValues.Set(key, val)
		}
		return strings.NewReader(urlValues.Encode()), nil
	case reflect.String:
		if header[common.HeaderContentType] != common.ContentTypeJSON.String() &&
			header[common.HeaderContentType] != common.ContentTypeText.String() {
			return nil, common.ErrInvalidContentTypeHeader
		}
		bodyStr, ok := body.(string)
		if !ok {
			return nil, common.ErrInvalidCastType
		}
		return strings.NewReader(bodyStr), nil
	default:
		return nil, nil
	}
}
