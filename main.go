package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Name string      `json:"name"`
	Date interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	// fmt.Println("Hello from go")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello from goserver")
	socket, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		msgType, msg, err := socket.ReadMessage()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(msg))

		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}
