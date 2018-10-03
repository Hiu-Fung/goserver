package main

import (
	"fmt"

	r "gopkg.in/gorethink/gorethink.v4"
)

type User struct {
	Id   string `gorething:"id,omitempty"`
	Name string `gorething:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
		//video 040 REQL in Go
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	user := User{
		Name: "anonymous",
	}

	// Note: doesn't return reponse, so better to use RunWrite()
	// err := r.Table("user")
	// 	.Insert(user)
	// 	.Exec(session)

	response, err := r.Table("user").
		Insert(user).
		RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}
	user2 := User{
		Name: "John Doe",
	}
	response2, _ := r.Table("user").
		Get("5d327257-79af-4bad-b0d5-e2adfd627f5e").
		Update(user2).
		RunWrite(session)

	fmt.Printf("%#v\n", response)
	fmt.Printf("%#v\n", response2)
}
