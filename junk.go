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
	fmt.Printf("%#v\n", response)
}
