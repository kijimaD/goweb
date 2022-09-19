// ユーザを認証し、クライアントにクッキーを返すハンドラ
func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email")) // ユーザ検索して構造体Userを返す
	if user.Password == data.Encrypt(r.PostFormValue("password")) { // データベース内に保存されているパスワードが、ハンドラに渡された暗号化パスワードと同じであることをチェックする
		session := user.CreateSession() // 取得したユーザのセッションを作成
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
