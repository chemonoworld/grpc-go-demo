package main

import (
	"github.com/chemonoworld/grpc-go-demo/client"
	"github.com/chemonoworld/grpc-go-demo/server"
)

func main() {

	//client.Start()
	go client.Start()
	go server.Start()

	for {

	}
}
