package main

import (
	"context"
	"github.com/giovanesantossilva/grpc/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main()  {
	connection, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)
	AddUserVerbose(client)
	AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient)  {
	req := &pb.User{
		Id: "123",
		Name: "Giovane",
		Email: "giovane1999@gmail.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}
	log.Printf("User addeded %v\n", res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id: "123",
		Name: "Giovane",
		Email: "giovane1999@gmail.com",
	}
	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		log.Println("Status user:", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient)  {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Giovane",
			Email: "giovane@gmail.com",
		},
		{
			Id:    "2",
			Name:  "Gustavo",
			Email: "gustavo@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receving response: %v", err)
	}

	log.Printf("Users addeded %v\n", res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Giovane",
			Email: "giovane@gmail.com",
		},
		{
			Id:    "2",
			Name:  "Gustavo",
			Email: "gustavo@gmail.com",
		},
	}

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			log.Printf("Sending user %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Could not receive the msg: %v", err)
			}
			log.Printf("User %v status %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}