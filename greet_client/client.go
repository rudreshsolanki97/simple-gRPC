package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/rudreshsolanki97/simple-gRPC/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("YO I am a CLient!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to host : %v", err)
	}
	c := greetpb.NewGreetServiceClient(cc)

	defer cc.Close()

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDiStreaming(c)
	// doBiDiStreamingCall(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	// fmt.Printf("Client created: %f", c)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rudresh",
			LastName:  "Solanki",
		},
	}
	if res, err := c.Greet(context.Background(), req); err != nil {
		log.Fatalf("error while calling Greet function: %v", err)
	} else {
		fmt.Printf("Response: %s \n", res.GetResult())
	}
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rudresh",
			LastName:  "Solanki",
		},
	}
	if resStream, err := c.GreetManyTimes(context.Background(), req); err != nil {
		log.Fatalf("error while calling Greet function: %v", err)
	} else {
		for {
			msg, err := resStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving the response from server: %v", err)
			}
			fmt.Printf("Response from stream: %s \n", msg.GetResult())
		}
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	firstNames := [3]string{"Raj", "Rahul", "Rohan"}
	lastNames := [3]string{"Shah", "Naik", "Desai"}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error in calling LongGreet service: %v", err)
	}
	for i := 0; i < 3; i++ {
		stream.Send(&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: firstNames[i],
				LastName:  lastNames[i],
			},
		})
	}
	res, newErr := stream.CloseAndRecv()
	if newErr != nil {
		log.Fatalf("Error in closing the client stream: %v", newErr)
	} else {
		fmt.Printf("Response: %s", res.GetResult())
	}
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	stream, err := c.BiDiGreet(context.Background())
	if err != nil {
		log.Fatalf("Error starting a client stream: %v", err)
	}

	firstNames := [3]string{"Raj", "Rahul", "Rohan"}
	lastNames := [3]string{"Shah", "Naik", "Desai"}

	waitc := make(chan struct{})
	// send a bunch of messages (go routines)
	go func() {
		// send a bunch of messages
		for i := 0; i < 3; i++ {
			stream.Send(&greetpb.ManyRequest{
				Greeting: &greetpb.Greeting{
					FirstName: firstNames[i],
					LastName:  lastNames[i],
				},
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	go func() {
		// receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving the response: %v", err)
				break
			}
			fmt.Printf("response: %s\n", res.GetResult())
		}
		close(waitc)
	}()
	<-waitc
}
