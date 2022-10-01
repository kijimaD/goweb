// ハンドラのチェイン

package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// Handler を受け取って Handlerを返す
func log(h http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc (func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("protect...")
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
}
