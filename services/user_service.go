package services

import (
	"context"

	"github.com/aleszilagyi/golang-grpc/config/logger"
	"github.com/aleszilagyi/golang-grpc/pb"
)

var log = logger.NewLogger()

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {

	log.Infof("Trying to insert to db: {%s}", request)
	log.Info("User inserted with success")

	return &pb.User{
		Id:    "123",
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}
