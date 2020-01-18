package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

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

	// Create a new gRPC server
	gRPCServer := grpc.NewServer()

	log.Printf("Search deployed on: %s\n", config.SearchServicePort)

	// Initialize new firestore client
	firestoreClient := internal.NewFirestoreClient()

	model.RegisterSearchServer(gRPCServer, &controllers.SearchServer{FirestoreClient: firestoreClient})

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
