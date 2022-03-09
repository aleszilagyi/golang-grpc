package main

import (
	"fmt"
	"net"

	"github.com/aleszilagyi/golang-grpc/config/logger"
	"github.com/aleszilagyi/golang-grpc/pb"
	"github.com/aleszilagyi/golang-grpc/resources"
	"github.com/aleszilagyi/golang-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	env *resources.AppEnv
	log  *logger.StandardLogger
)

func init() {
	log = logger.NewLogger()
	env = resources.GetConf()
}

func main() {
	address := fmt.Sprintf("%s:%s", env.AppConfig.Hostname, env.AppConfig.Port)

	log.Info("starting application")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("could not connect: %v", err)
	}
	log.Infof("server started listening to %v", address)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Errorf("could not serve: %v", err)
	}
}
