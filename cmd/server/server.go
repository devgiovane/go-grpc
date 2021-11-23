package main

import (
	"github.com/giovanesantossilva/grpc/pb"
	"github.com/giovanesantossilva/grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main()  {
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cold not server: %v", err)
	}
}
