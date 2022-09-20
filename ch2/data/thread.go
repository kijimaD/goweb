// このパッケージ内ものをあとで使うときは、data.Threadとしなければならない*
package data

import (
	"time"
)

// DDL(データ定義言語)に対応するものであり、threadsというRDBテーブルを作成するのに使われる
type Thread struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	// データベースに接続する
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread()	// 構造体を作成する
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th) // 行のデータをスライスに追加していく
	}
	rows.Close()
	return
}

// スレッド数のcountを返すメソッド
// データ構造に対する関数やメソッド(User, Session, Thread, Post)を提供することで、ハンドラ関数内でデータベースに直接アクセスすることを防ぐデータ層を作成できる
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close() // エラーがあったらここまで到達しDBとの接続を終了
	return
}
