package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// hello worldを返すエンドポイントの作成
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// サーバーの起動
	http.ListenAndServe(":8080", nil)

}
