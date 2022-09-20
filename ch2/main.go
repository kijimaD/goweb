// マルチプレクサ: リクエストをハンドラにリダイレクトするコード

package main

import (
	"net/http"
	"html/template"
)

func main() {
	mux := http.NewServeMux() // デフォルトのマルチプレクサ
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files)) // マルチプレクサはハンドラへのリクエストに使えるだけでなく、静的なファイルの返送にも使える。
	// /static/で始まるすべてのリクエストURLについて、URLから文字列/staticを取り去り、ディレクトリpublicを起点として、残った文字列を名前にもつファイルを探す。
	// /static/example.txt の場合は、./example.txt を探す。ここでいうルートディレクトリはpublicなので、リポジトリ的には、./public/example.txt を探す。

	mux.HandleFunc("/", index) // ルートURLをハンドラ関数にリダイレクトする

	// マルチプレクサをサーバに付加
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

// ハンドラ関数は第1引数にResponseWriterを取り、第2引数にRequestへのポインタを取るGoの関数にすぎない
// HTMLを生成してResponseWriterに書き出す
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
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
