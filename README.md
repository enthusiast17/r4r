# r4r

Router and handler wrapper for "net/http" from standard library. This is inspired by [Express.js](https://github.com/expressjs/express).

### Installation

```
go get github.com/enthusiast17/r4r
```

### Usage

```Go
package main

import "r4r"

var PORT int = 5000
func thisIsMiddleware(w http.ResponseWriter, r *http.Request) *r4r.Error { ... }

func main() {
  router := r4r.NewRouter()
  
  router.Get("/hello_world", r4r.withMiddlewares(thisIsMiddleware), func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
  })
  
  r4r.Use("", nil, router)
  
  _ := http.ListenAndServe(PORT, nil)
}

```
