// ハンドラ関数によるリクエストの処理。複数のハンドラ関数を使った例と比較する

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello) // ハンドラ関数HandleFunc。関数helloをハンドラに変換して、そのハンドラをDefaultServeMuxに登録する。ハンドラ関数とは、単にハンドラを作成するための便利な手段
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
