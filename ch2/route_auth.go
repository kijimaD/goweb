package main

import (
	"example/data"
	"net/http"
)

// GET /login
// Show the login page
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := data.User {
		Name: request.PostFormValue("name"),
		Email: request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// ユーザを認証し、クライアントにクッキーを返すハンドラ
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email")) // ユーザ検索して構造体Userを返す
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) { // データベース内に保存されているパスワードが、ハンドラに渡された暗号化パスワードと同じであることをチェックする
		session, err := user.CreateSession() // 取得したユーザのセッションを作成
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{ // セッション情報をクッキーに書き込む
			Name: "_cookie",
			Value: session.Uuid, // 最も重要な値。これをブラウザに保存したい。IDなのであとから引ける
			HttpOnly: true, // クッキーへのアクセスをHTTPとHTTPSにのみ許可。jsで書き換えられないように
		}
		http.SetCookie(w, &cookie) // レスポンスヘッダーにクッキを書く。ユーザがログインしたら、その後のリクエストではユーザがすでにログインしているということが示されている必要がある。クッキーはクライアントに送信され、ブラウザに保存される
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
