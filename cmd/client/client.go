package main

import (
	"context"
	"fmt"

	"github.com/aleszilagyi/golang-grpc/config/logger"
	"github.com/aleszilagyi/golang-grpc/pb"
	"github.com/aleszilagyi/golang-grpc/resources"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	env *resources.AppEnv
	log *logger.StandardLogger
)

func init() {
	env = resources.GetConf()
	log = logger.NewLogger()
}

func main() {
	connectionString := fmt.Sprintf("%s:%s", env.AppConfig.Hostname, env.AppConfig.Port)

	connection, err := grpc.Dial(connectionString, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("could not connect to gRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Ale",
		Email: "ale@ale.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Errorf("error when trying to make gRPC request: %v", err)
	}

	log.Infof("{%s}", res)
}
