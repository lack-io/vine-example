package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine/service"
	"github.com/lack-io/vine/service/client"
)

func main() {
	srv := service.NewService()
	srv.Init()
	c := srv.Client()

	request := c.NewRequest("greeter", "Greeter.Hello", "john", client.WithContentType("application/json"))
	var response string

	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
