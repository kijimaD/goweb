// ハンドラを作成してサーバに割り当てる。マルチプレクサを使ってないので、リクエストを特定のハンドラに転送するURLのマッチング処理はない。サーバに届くすべてのリクエストがこのハンドラに行く。

package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &handler, // 普通は指定しない。指定しない倍DefaultServeMuxがマルチプレクサとして使われる
	}
	server.ListenAndServe()
}
