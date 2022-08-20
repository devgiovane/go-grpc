package services

import (
	"context"
	"github.com/giovanesantossilva/grpc/pb"
	"io"
	"log"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	log.Printf("Addeded user: %v\n", req)
	return &pb.User{Id: req.GetId(), Name: req.GetName(), Email: req.GetEmail()}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer)  error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User: &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User: &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserted",
		User: &pb.User{Id: "123", Name: req.GetName(), Email: req.GetEmail()},
	})

	log.Printf("Addeded user: %v\n", req)
	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	var users []*pb.User

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		users = append(users, &pb.User{
			Id: req.GetId(),
			Name: req.GetName(),
			Email: req.GetEmail(),
		})

		log.Printf("Adding %v\n", req)
	}
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User: req,
		})
		if err != nil {
			log.Fatalf("Error sending stream to the cleint: %v", err)
		}
	}
}