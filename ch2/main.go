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

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

// ハンドラ関数は第1引数にResponseWriterを取り、第2引数にRequestへのポインタを取るGoの関数にすぎない
// HTMLを生成してResponseWriterに書き出す
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads(); if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
