// 複数のハンドラによるリクエストの処理。
// ハンドラ関数でリクエスト処理を行う
// 自作の構造体にServeHTTPメソッドを実装してハンドラにして、パスと紐付ける

package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", &hello) // 関数HandleはURLに対応するハンドラへのポインタを登録する
	http.Handle("/world", &world)

	server.ListenAndServe()
}
