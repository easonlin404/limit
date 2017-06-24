# Limit Gin's middleware

[![Travis branch](https://img.shields.io/travis/easonlin404/limit/master.svg)](https://travis-ci.org/easonlin404/limit)
[![Codecov branch](https://img.shields.io/codecov/c/github/easonlin404/limit/master.svg)](https://codecov.io/gh/easonlin404/limit)
[![Go Report Card](https://goreportcard.com/badge/github.com/easonlin404/limit)](https://goreportcard.com/report/github.com/easonlin404/limit)
[![GoDoc](https://godoc.org/github.com/easonlin404/limit?status.svg)](https://godoc.org/github.com/easonlin404/limit)
 
Gin middleware to limit the number of current requests.

## Usage

### Start using it

Download and install it:

```sh
$ go get github.com/easonlin404/limit
```

Import it in your code:

```go
import "github.com/easonlin404/limit"
```

### Canonical example:

```go
package main

import (
	"github.com/easonlin404/limit"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(limit.Limit(200)) // limit the number of current requests

	r.GET("/", func(c *gin.Context) {
		// your code
	})

	r.Run()
}
```