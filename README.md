# Status Snatcher
Package to fetch the status code from `net/http`'s [ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter)
## Install
```
go install github.com/maheshwar-mr/statussnatcher
```
## Usage
Sample setup for a logging middleware.
```go
package middleware

import (
    "log"
    "net/http"
    "time"

    "github.com/maheshwar-mr/statussnatcher"
)

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        start := time.Now()
        ss := statussnatcher.New(w) //Create a new instance
        next.ServeHTTP(ss, r) //Calling ServeHTTP() of the handler
        log.Println(ss.StatusCode, r.Method, r.URL.Path, time.Since(start))
    })
}
```
