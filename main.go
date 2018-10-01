package main

import (
	"net/http"
)

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := NewRouter()
	// router.Handle("channel add", addChannel)
	// http.HandleFunc("/", handler)
	http.Handle("/", router)
	http.ListenAndServe(":4001", nil)
}
