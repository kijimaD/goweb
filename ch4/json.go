package main

import(
	"encoding/json"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "Sau Sheong",
		Threads: []string{"1番目", "2番目", "3番目"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}


// curl -i 127.0.0.1:8080/json


// HTTP/1.1 200 OK
// Content-Type: application/json
// Date: Sat, 01 Oct 2022 05:49:14 GMT
// Content-Length: 63

// {"User":"Sau Sheong","Threads":["1番目","2番目","3番目"]}
