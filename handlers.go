package main

import (
	"github.com/mitchellh/mapstructure"
	r "gopkg.in/gorethink/gorethink.v4"
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	// var message Message
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}

	go func() {
		err = r.Table("channel").
			Insert("channel").
			Exec(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
		}
	}()
	// channel.Id = "ABC123"
	// message.Name = "channel add"
	// message.Data = channel
	// client.send <- message
}
