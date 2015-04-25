# go-httpslash [![Build Status](https://travis-ci.org/frozzare/go-httpslash.svg?branch=master)](https://travis-ci.org/frozzare/go-httpslash)

Enforce trailing slash or no trailing slash.

View the [docs](http://godoc.org/github.com/frozzare/go-httpslash).

## Installation

```
$ go get github.com/frozzare/go-httpslash
```

## Example

Example using `http.ServerMux()`

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/frozzare/go-httpslash"
)

func main() {
	mux := http.NewServeMux()
    // /hello will redirect to /hello/
	mux.Handle("/", httpslash.TrailingSlash(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL, %s!", r.URL.String())
	})))
	http.ListenAndServe(":3000", mux)
}
```

Example using [Goji](http://github.com/zenazn/goji)

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/frozzare/go-httpslash"
	"github.com/zenazn/goji/web"
)

func main() {
	mux := web.New()
	mux.Use(httpslash.TrailingSlash)
    // /hello will redirect to /hello/
	mux.Get("/hello/", func(c web.C, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL, %s!", r.URL.String())
	})
	http.ListenAndServe(":3000", mux)
}
```

# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
