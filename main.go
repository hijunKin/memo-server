package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var htmlStr string

func main() {
  	fmt.Println("start")

	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	htmlStr = string(data)

	http.HandleFunc("/", showHTML)
	http.HandleFunc("/add_memo", addMemo)
	http.HandleFunc("/list_memos", listMemos)

	http.ListenAndServe(":8080", nil) 
}

//curl localhost:8080/
func showHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, htmlStr)
}

type Memo struct {
	ID string
	Title string 
	Body string
	CreatedAd time.Time
	UpdatedAt time.Time
}

var memos = map[string]*Memo{}

//curl http://localhost:8080/add_memo -X POST -H "Content-Type: application/json" -d '{"ID": "tio"}'
//curl localhost:8080/add_memo -X POST -H "Content-Type: application/json" -d '{"ID": "1111", "Title": "myTitle", "Body": "myBody", "CreatedAt": "2022-01-01T10:00:00Z", "UpdatedAt": "2022-01-01T11:00:00Z"}'
func addMemo(w http.ResponseWriter, r *http.Request) {
	m := &Memo{}

	//Httpリクエストで送信されてきたHttpRequestBody(JSON形式)をMemo構造体にセットしている。
	if err := json.NewDecoder(r.Body).Decode(m); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	memos[m.ID] = m

	fmt.Fprintln(w, len(memos))
}

//curl localhost:8080/list_memos
func listMemos(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(memos)
	if err != nil {
		fmt.Fprintln(w, "error:"+err.Error())
		return 
	}

	fmt.Fprintln(w, string(b))
}