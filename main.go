package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
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
		// msgType, msg, err := socket.ReadMessage()

		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		var inMessage Message
		var outMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("%#v\n", inMessage) //https://jsbin.com/rixicuzivu/edit?js,console,output

		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage = Message{"error", err}
				if err := socket.WriteJSON(outMessage); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			err := subscribeChannel()
		}
		// fmt.Println(string(msg))

		// if err = socket.WriteMessage(msgType, msg); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
}

func addChannel(data interface{}) error {
	var channel Channel

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}

	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	return nil
}

func subscribeChannel() {
	//TODO: rethinkDB Query / changefeed
	for {
		time.Sleep(time.Second * 1)
	}
}
