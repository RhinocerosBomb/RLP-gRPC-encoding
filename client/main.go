package main

import (
	"awesomeProject1/codec"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"os"
	"time"
)


const (
	address     = "localhost:8080"
	defaultName = "world"
)

func init()  {
	encoding.RegisterCodec(&codec.RLPEncoder{})
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.CallContentSubtype("rlp"))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
