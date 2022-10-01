package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}

// map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8] Accept-Encoding:[gzip, deflate, br] Accept-Language:[ja,en-US;q=0.5] Connection:[keep-alive] Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (X11; Linux x86_64; rv:104.0) Gecko/20100101 Firefox/104.0]]


map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8]
Accept-Encoding:[gzip, deflate, br] Accept-Language:[ja,en-US;q=0.5]
Connection:[keep-alive]
Sec-Fetch-Dest:[document]
Sec-Fetch-Mode:[navigate]
Sec-Fetch-Site:[none]
Sec-Fetch-User:[?1]
Upgrade-Insecure-Requests:[1]
User-Agent:[Mozilla/5.0 (X11; Linux x86_64; rv:104.0) Gecko/20100101 Firefox/104.0]]
