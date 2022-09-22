package main

import (
	"example/data"
	"net/http"
)

// ハンドラ関数は第1引数にResponseWriterを取り、第2引数にRequestへのポインタを取るGoの関数にすぎない
// HTMLを生成してResponseWriterに書き出す

// クライアントに返信するHTMLの生成
func index(writer http.ResponseWriter, request *http.Request){
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}

// GET /err?msg=
// show the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
