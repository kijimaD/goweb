package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// HandlerFuncを受け取り、HandlerFuncを返す
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // 無名関数
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name() // 引数の関数名を取得する -- 例: main.hello
		fmt.Println("Handler function called - " + name) // 引数のHandlerFuncの名前を出力してから、引数の関数を呼び出す。つまりlog -> helloという処理
		h(w, r)
	}
}

func protect(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protect code...")
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", protect(log(hello))) // 関数helloを関数logに送り込んでいる。logとhelloをチェインさせている
	server.ListenAndServe()
}
