package middleware

import (
	"fmt"
	"net/http"
	"time"
)
func AllInfoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start:= time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s, %s %s\r", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}