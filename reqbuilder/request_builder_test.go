package reqbuilder

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"unicode"
)

var requestBuilder = NewRequestBuilder().SetMethodGet().SetBaseURL("www.foo.com")

func TestNewRequestBuilder(t *testing.T) {
	assert.Implements(t, (*RequestBuilder)(nil), NewRequestBuilder())
}

func Test_requestBuilderImpl_SetMethodGet(t *testing.T) {
	request, _ := requestBuilder.SetMethodGet().Build()

	assert.Equal(t, http.MethodGet, request.Method)
}

func Test_requestBuilderImpl_SetMethodPut(t *testing.T) {
	request, _ := requestBuilder.SetMethodPut().Build()

	assert.Equal(t, http.MethodPut, request.Method)
}

func Test_requestBuilderImpl_SetMethodPost(t *testing.T) {
	request, _ := requestBuilder.SetMethodPost().Build()

	assert.Equal(t, http.MethodPost, request.Method)
}

func Test_requestBuilderImpl_SetMethodPatch(t *testing.T) {
	request, _ := requestBuilder.SetMethodPatch().Build()

	assert.Equal(t, http.MethodPatch, request.Method)
}

func Test_requestBuilderImpl_SetMethodDelete(t *testing.T) {
	request, _ := requestBuilder.SetMethodDelete().Build()

	assert.Equal(t, http.MethodDelete, request.Method)
}

func Test_requestBuilderImpl_SetMethodOptions(t *testing.T) {
	request, _ := requestBuilder.SetMethodOptions().Build()

	assert.Equal(t, http.MethodOptions, request.Method)
}

func Test_requestBuilderImpl_SetBaseURL_SetPath(t *testing.T) {
	request, _ := requestBuilder.SetPath("/foo").Build()

	assert.Equal(t, "www.foo.com/foo", request.URL.String())
}

func Test_requestBuilderImpl_SetBody(t *testing.T) {
	type BodyRequest struct {
		Message string `json:"message"`
	}
	body := &BodyRequest { Message: "foo" }
	requestWithBody, _ := requestBuilder.SetBody(body).Build()

	var requestBody map[string]string
	_ = json.NewDecoder(requestWithBody.Body).Decode(&requestBody)

	assert.Equal(t,"foo", requestBody["message"])
}

func Test_requestBuilderImpl_SetBody_error(t *testing.T) {
	_, err := requestBuilder.SetBody(make(chan int)).Build()

	assert.Equal(t, "json: unsupported type: chan int", err.Error())
}

func Test_requestBuilderImpl_SetBody_empty(t *testing.T) {
	request, _ := requestBuilder.Build()

	assert.Equal(t, nil, request.Body)
}

func Test_requestBuilderImpl_AddHeader(t *testing.T) {
	key := "foo-key"
	expectedValue := "foo-value"
	request, _ := requestBuilder.AddHeader(key, expectedValue).Build()

	assert.Equal(t, 1, len(request.Header))
	assert.Equal(t, expectedValue, request.Header.Get(key))
}

func Test_requestBuilderImpl_AddQueryParam(t *testing.T) {
	key := "foo-key"
	expectedValue := "foo value"
	request, _ := requestBuilder.AddQueryParameter(key, expectedValue).Build()

	assert.Equal(t, "foo-key=foo+value", request.URL.RawQuery)
}

func Test_requestBuilderImpl_Build(t *testing.T) {
	request, err := requestBuilder.Build()

	assert.Nil(t, err)
	assert.IsType(t, (*http.Request)(nil), request)
}

func Test_apiRequestBuilderImpl_Build_errorMethodNotDefined(t *testing.T) {
	request, err := NewRequestBuilder().Build()

	assert.Nil(t, request)
	assert.Equal(t, "HTTP method is not defined", err.Error())
}

func Test_apiRequestBuilderImpl_Build_errorBaseURLNotDefined(t *testing.T) {
	request, err := NewRequestBuilder().SetMethodGet().Build()

	assert.Nil(t, request)
	assert.Equal(t, "base URL is not define", err.Error())
}

func Test_apiRequestBuilderImpl_Build_error_net_URL(t *testing.T) {
	request, err := NewRequestBuilder().SetMethodGet().SetBaseURL(string(unicode.MaxASCII)).Build()

	assert.Nil(t, request)
	assert.Equal(t, "parse \"\\u007f\": net/url: invalid control character in URL", err.Error())
}