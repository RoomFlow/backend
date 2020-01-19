package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/RoomFlow/backend/internal/config"
	internal "github.com/RoomFlow/backend/internal/helpers"
	model "github.com/RoomFlow/backend/internal/proto/search"
	"github.com/RoomFlow/backend/services/search/controllers"
)

func main() {
	lis, err := net.Listen("tcp", config.SearchServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create server certificates.
	creds, err := credentials.NewServerTLSFromFile(config.SSLCertPath, config.SSLKeyPath)
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}

	// Create a new server using the created credentials.
	gRPCServer := grpc.NewServer(grpc.Creds(creds))

	// Initialize new firestore client
	firestoreClient := internal.NewFirestoreClient()

	model.RegisterSearchServer(gRPCServer, &controllers.SearchServer{FirestoreClient: firestoreClient})

	log.Printf("Search deployed on: %s\n", config.SearchServicePort)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
