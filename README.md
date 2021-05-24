[![Release](https://img.shields.io/github/v/release/juanenmellare/gorequestbuilder.svg?style=flat-square)](https://github.com/juanenmellare/gorequestbuilder/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/juanenmellare/gorequestbuilder.svg)](https://pkg.go.dev/github.com/juanenmellare/gorequestbuilder)
[![CircleCI](https://circleci.com/gh/juanenmellare/gorequestbuilder.svg?style=shield)](https://circleci.com/gh/juanenmellare/gorequestbuilder)
[![codecov](https://codecov.io/gh/juanenmellare/gorequestbuilder/branch/main/graph/badge.svg?token=ZCRF68IC8Z)](https://codecov.io/gh/juanenmellare/gorequestbuilder)

# <img width="60px" align="center" src="https://miro.medium.com/fit/c/262/262/1*yh90bW8jL4f8pOTZTvbzqw.png">Go Request Builder
A simple request (http.Request) builder for Golang.

## Import

```go
import "github.com/juanenmellare/gorequestbuilder"
```

## Quick Start
In every request buillder set the HTTP method and the base URL otherwise it'll return a nil request and an error.
```go
request, err := gorequestbuilder.NewRequestBuilder().
    SetMethodGet().
    SetBaseURL("https://en.wikipedia.org").
    SetPath("/wiki/Go_(programming_language)").
    Build()
 ```
 
 ## Headers
```go
var requestBuilder = gorequestbuilder.NewRequestBuilder()

requestBuilder.AddHeader("Authorization", "Basic R29sYW5nIERldmVsb3Blcg==")
```

 ## Query Parameters
```go
var requestBuilder = gorequestbuilder.NewRequestBuilder()

requestBuilder.AddQueryParameter("query", "language=golang")
```

## Body
```go
type Body struct {
	Message string `json:"message"`
}
	
body := &Body{
	Message: "Golang is awesome!",
}

var requestBuilder = gorequestbuilder.NewRequestBuilder()

requestBuilder.SetMethodPost().SetBody(body)
```

## Build
```go
var requestBuilder = gorequestbuilder.NewRequestBuilder()
.
.
.
request, err = requestBuilder.Build()
```
