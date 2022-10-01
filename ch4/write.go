package main

import (
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
<html>`
	w.Write([]byte(str))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	server.ListenAndServe()
}

// curl -i 127.0.0.1:8080/write


// HTTP/1.1 200 OK
// Date: Sat, 01 Oct 2022 05:19:03 GMT
// Content-Length: 94
// Content-Type: text/html; charset=utf-8

// <html>
// <head><title>Go Web Programming</title></head>
// <body><h1>Hello World</h1></body>


// コンテンツタイプが自動で検知されている
