// マルチプレクサ: リクエストをハンドラにリダイレクトするコード

package main

import (
	"net/http"
	"html/template"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	mux := http.NewServeMux() // デフォルトのマルチプレクサ
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files)) // マルチプレクサはハンドラへのリクエストに使えるだけでなく、静的なファイルの返送にも使える。
	// /static/で始まるすべてのリクエストURLについて、URLから文字列/staticを取り去り、ディレクトリpublicを起点として、残った文字列を名前にもつファイルを探す。
	// /static/example.txt の場合は、./example.txt を探す。ここでいうルートディレクトリはpublicなので、リポジトリ的には、./public/example.txt を探す。

	// index
	mux.HandleFunc("/", index) // ルートURLをハンドラ関数にリダイレクトする
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thrad/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// サーバを開始、マルチプレクサをサーバに付加する
	// マルチプレクサとは、複数の信号を受け、それらを選択したりまとめたりして一つの信号として出力する装置や素子、機構などのこと
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

// Webアプリケーションの流れ

// 1. クライアントがサーバにリクエストを送る
// 2. リクエストはマルチプレクサが受信し、適切なハンドラにリダイレクトする
// 3. ハンドラがリクエストを処理する
// 4. データが必要な場合は、データベース内のデータをモデル化
// 5. モデルはデータベースと接続する。これはデータ構造の関数やメソッドによって発動されて行われる
// 6. 処理が完了すると、ハンドラはテンプレートエンジンを起動し、場合によってモデルからデータを送信する
// 7. テンプレートエンジンがテンプレートファイルを解析しテンプレートを作成する。テンプレートはデータと組み合わされてHTMLとなる
// 8. 生成されたHTMLはレスポンスの一部としてクライアントに返送される
