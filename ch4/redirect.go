package main

import (
	"net/http"
)

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)	// WriteHeaderは呼び出された後にヘッダが変更されるのを防ぐので、ステータスコードを書き込む前にLocationヘッダを追加しなければならない
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/redirect", headerExample)
	server.ListenAndServe()
}


// curl -i 127.0.0.1:8080/redirect

// HTTP/1.1 302 Found
// Location: http://google.com
// Date: Sat, 01 Oct 2022 05:37:47 GMT
// Content-Length: 0
