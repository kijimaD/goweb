package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	// Responseから情報を取り出してHTTPレスポンスを生成し、そのレスポンスがResponseWriterを介して返信される
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler) // ルートURLが呼び出されたとき、定義されたハンドラを起動する
	http.ListenAndServe(":8080", nil) // ポート8080を監視するようにサーバを起動
}

// アプリケーションサーバもパッケージnet/httpが提供し、一緒にコンパイルされる(別でアプリケーションサーバが必要なRubyなどとは異なる)。単独で配置可能なスタンドアローンのWebアプリが生成される。
