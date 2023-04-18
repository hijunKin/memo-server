package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var htmlStr string

func main() {
  fmt.Println("start")

  data, err := os.ReadFile("index.html")
  if err != nil {
    log.Fatal(err)
  }

  htmlStr = string(data)

	http.HandleFunc("/", handler) //ハンドラを登録---(1)

	http.ListenAndServe(":8080", nil) //サーバーを起動---(2)
}

//HelloHandler サーバーの処理内容を記述---(3)
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, count)
	fmt.Fprintln(w, htmlStr)
}