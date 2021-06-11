package gorequestbuilder

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// RequestBuilder is the main interface with some methods to build and handle a request.
type RequestBuilder interface {
	SetMethodGet() RequestBuilder
	SetMethodPut() RequestBuilder
	SetMethodPost() RequestBuilder
	SetMethodPatch() RequestBuilder
	SetMethodDelete() RequestBuilder
	SetMethodOptions() RequestBuilder
	SetBaseURL(baseURL string) RequestBuilder
	SetPath(path string) RequestBuilder
	SetBody(body interface{}) RequestBuilder
	AddHeader(key, value string) RequestBuilder
	AddQueryParameter(key, value string) RequestBuilder
	Build() (*http.Request, error)
}

type requestBuilderImpl struct {
	headers         map[string]string
	queryParameters map[string]string
	body            interface{}
	baseURL         string
	path            string
	httpMethod      string
}

// NewRequestBuilder is the constructor of an empty RequestBuilder.
func NewRequestBuilder() RequestBuilder {
	return &requestBuilderImpl{
		headers:         make(map[string]string),
		queryParameters: make(map[string]string),
	}
}

func (r requestBuilderImpl) SetMethodGet() RequestBuilder {
	r.httpMethod = http.MethodGet
	return &r
}

func (r requestBuilderImpl) SetMethodPut() RequestBuilder {
	r.httpMethod = http.MethodPut
	return &r
}

func (r requestBuilderImpl) SetMethodPost() RequestBuilder {
	r.httpMethod = http.MethodPost
	return &r
}

func (r requestBuilderImpl) SetMethodPatch() RequestBuilder {
	r.httpMethod = http.MethodPatch
	return &r
}

func (r requestBuilderImpl) SetMethodDelete() RequestBuilder {
	r.httpMethod = http.MethodDelete
	return &r
}

func (r requestBuilderImpl) SetMethodOptions() RequestBuilder {
	r.httpMethod = http.MethodOptions
	return &r
}

func (r requestBuilderImpl) SetBaseURL(baseURL string) RequestBuilder {
	r.baseURL = baseURL
	return &r
}

func (r requestBuilderImpl) SetPath(path string) RequestBuilder {
	r.path = path
	return &r
}

func (r requestBuilderImpl) SetBody(body interface{}) RequestBuilder {
	r.body = body
	return &r
}

func (r requestBuilderImpl) AddHeader(key, value string) RequestBuilder {
	r.headers[key] = value
	return &r
}

func (r requestBuilderImpl) AddQueryParameter(key, value string) RequestBuilder {
	r.queryParameters[key] = value
	return &r
}

func (r requestBuilderImpl) Build() (*http.Request, error) {
	if err := validateEmptyStringField(r.httpMethod, "HTTP method"); err != nil {
		return nil, err
	}
	if err := validateEmptyStringField(r.baseURL, "base URL"); err != nil {
		return nil, err
	}

	reader, errReader := parseBodyJSONToReader(r.body)
	if errReader != nil {
		return nil, errReader
	}

	request, errRequest := http.NewRequest(r.httpMethod, r.baseURL+r.path, reader)
	if errRequest != nil {
		return nil, errRequest
	}

	setHeaders(request, r.headers)
	setQueryParameters(request, r.queryParameters)

	return request, nil
}

func validateEmptyStringField(fieldValue, errMessage string) error {
	if fieldValue == "" {
		return errors.New(errMessage + " is not defined")
	}

	return nil
}

func parseBodyJSONToReader(body interface{}) (io.Reader, error) {
	var reader io.Reader
	if body == nil {
		return reader, nil
	}

	jsonBody, err := json.MarshalIndent(body, "", " ")
	if err != nil {
		return nil, err
	}
	bytesReader := bytes.NewReader(jsonBody)

	return bytesReader, nil
}

func setHeaders(request *http.Request, headers map[string]string) {
	if len(headers) != 0 {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}
}

func setQueryParameters(request *http.Request, queryParameters map[string]string) {
	if len(queryParameters) != 0 {
		query := request.URL.Query()
		for key, value := range queryParameters {
			query.Set(key, value)
		}
		request.URL.RawQuery = query.Encode()
	}
}
