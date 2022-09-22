// 最も単純なWebサーバ

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("", nil)
}
