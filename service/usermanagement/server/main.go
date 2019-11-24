package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	userManagement "github.com/RoomFlow/backend/proto/usermanagement"
)

const (
	port = ":50051"
)

type Config struct {
	GRPCPort string
}


type userManagementServer struct {
	userManagement.UserManagementServer
}


// Attempts to login a user based on the passed on credentials.
func (s *userManagementServer) LoginUser(ctx context.Context, request *userManagement.LoginRequest) (*userManagement.LoginResponse, error) {
	// TODO: access the database and check for correctness of credentials
	return &userManagement.LoginResponse{Token: "zxyshgfaewgesf"}, nil
}

// Attempts to register a user based on the passed on credentials.
func (s *userManagementServer) RegisterUser(ctx context.Context, request *userManagement.RegisterRequest) (*userManagement.RegisterResponse, error) {
	// TODO: add the new user to the database.
	return &userManagement.RegisterResponse{Token: "zszvjnrigearvbgr"}, nil
}

func newServer() *userManagementServer {
	server := &userManagementServer{}
	return server
}

func main() {
	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.Parse()

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
			log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userManagement.RegisterUserManagementServer(grpcServer, &userManagementServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}