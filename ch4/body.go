package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}


// このコードを試すには適切なメッセージボディを設定したPOSTリクエストをサーバに送信する必要がある。GETリクエストにはメッセージボディがないため、このコードだと何もprintしない
// $ curl -id "first_name=sausheong_last_name=chang" 127.0.0.1:8080/body


// HTTP/1.1 200 OK
// Date: Sat, 01 Oct 2022 01:12:30 GMT
// Content-Length: 37
// Content-Type: text/plain; charset=utf-8

// first_name=sausheong_last_name=chang
