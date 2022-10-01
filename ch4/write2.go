package main

import (
	"fmt"
	"net/http"
)

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/writeheader", writeHeaderExample)
	server.ListenAndServe()
}


// curl -i 127.0.0.1:8080/writeheader


// HTTP/1.1 501 Not Implemented
// Date: Sat, 01 Oct 2022 05:28:30 GMT
// Content-Length: 46
// Content-Type: text/plain; charset=utf-8

// そのようなサービスはありません
