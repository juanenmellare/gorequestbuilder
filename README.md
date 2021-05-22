[![Release](https://img.shields.io/github/v/release/juanenmellare/gorequestbuilder.svg?style=flat-square)](https://github.com/juanenmellare/gorequestbuilder/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/juanenmellare/gorequestbuilder.svg)](https://pkg.go.dev/github.com/juanenmellare/gorequestbuilder)
[![CircleCI](https://circleci.com/gh/juanenmellare/gorequestbuilder.svg?style=shield)](https://circleci.com/gh/juanenmellare/gorequestbuilder)
[![codecov](https://codecov.io/gh/juanenmellare/gorequestbuilder/branch/main/graph/badge.svg?token=ZCRF68IC8Z)](https://codecov.io/gh/juanenmellare/gorequestbuilder)

<img align="right" width="140px" src="https://www.clipartmax.com/png/small/111-1112912_go-gopher-go-programming-language-logo.png">

# Go Request Builder
A simple request (http.Request) builder for Golang.

## Import

```go
import "github.com/juanenmellare/gorequestbuilder"
```

## Quick Start
```go
request, err := gorequestbuilder.NewRequestBuilder().
    SetMethodGet().
    SetBaseURL("https://golang.org/").
    SetPath().
    Build()
 ```
 
 ## Headers
```go
var requestBuilder = gorequestbuilder.NewRequestBuilder()
requestBuilder.AddHeader("Authorization", "Basic R29sYW5nIERldmVsb3Blcg==")
```


