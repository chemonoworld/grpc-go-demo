package client

import (
	"context"
	"fmt"
	"github.com/chemonoworld/grpc-go-demo/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Start() {
	// client에서 server로 연결
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	//defer conn.Close()
	client := proto.NewGreeterClient(conn)

	name := "Alice"
	for {

		response, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("Error calling SayHello: %v", err)
		}
		fmt.Printf("Response from server: %s\n", response.Message)
		time.Sleep(1 * time.Second)
	}
}
