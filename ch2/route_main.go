package main

import (
	"example/data"
	"net/http"
)

// ハンドラ関数は第1引数にResponseWriterを取り、第2引数にRequestへのポインタを取るGoの関数にすぎない
// HTMLを生成してResponseWriterに書き出す

// クライアントに返信するHTMLの生成
func index(writer http.ResponseWriter, request *http.Request){
	threads, err := data.Threads(); if err == nil {
		_, err := session(w, r) // ここではerrしか使わないので、構造体sessionは不要
		public_teml_files := []string{"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html"}
		private_tmpl_files := []string{"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html"}
		var templates *template.Template
		if err != nil {
			// テンプレートファイルを解析してテンプレートを作成する
			// 解析したあと、結果を関数Mustに渡さなければならない。エラーを検知するため。
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		}
		// 解析しておいたテンプレートを実行する
		// テンプレートファイルからコンテンツを取り出し、別のところから得られるデータと組み合わせて、最終的なHTMLコンテンツを生成する
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

// GET /err?msg=
// show the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URLQuery()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
