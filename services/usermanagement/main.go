package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/RoomFlow/backend/internal/config"
	model "github.com/RoomFlow/backend/internal/proto/usermanagement"
	"github.com/RoomFlow/backend/services/usermanagement/controllers"
)

func main() {
	lis, err := net.Listen("tcp", config.UsermanagementServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create server certificates.
	serverCert, err := credentials.NewServerTLSFromFile(config.SSLCertPath, config.SSLKeyPath)
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}

	// Create a new server using the created credentials.
	gRPCServer := grpc.NewServer(grpc.Creds(serverCert))

	// Register the created user management server.
	model.RegisterUserManagementServer(gRPCServer, &controllers.UsermanagementServer{})

	log.Printf("Usermanagement deployed on %s\n", config.UsermanagementServicePort)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
