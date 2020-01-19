package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"

	model_search "github.com/RoomFlow/backend/internal/proto/search"
	model_usermanagement "github.com/RoomFlow/backend/internal/proto/usermanagement"
)

// CreateGateway creates the gateway to access other microservices.
func CreateGateway(ctx context.Context, muxOptions ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(muxOptions...)

	// Create user management credentials.
	searchCredentials, err := credentials.NewClientTLSFromFile("internal/certs/app.crt", "")
	if err != nil {
		log.Fatalf("Error in creating server credentials. %v", err)
		return nil, err
	}
	// Create user management dial options using the created credentials.
	searchDialOptions := []grpc.DialOption{grpc.WithTransportCredentials(searchCredentials)}

	// Register the user management handler from endpoint using the created dial options.
	err = model_search.RegisterSearchHandlerFromEndpoint(ctx, mux, "localhost:10001", searchDialOptions)
	if err != nil {
		log.Fatalf("Error in registering end point. %v", err)
		return nil, err
	}

	// Create user management credentials.
	userManagementCredentials, err := credentials.NewClientTLSFromFile("internal/certs/app.crt", "")
	if err != nil {
		log.Fatalf("Error in creating server credentials. %v", err)
		return nil, err
	}
	// Create user management dial options using the created credentials.
	userManagementDialOptions := []grpc.DialOption{grpc.WithTransportCredentials(userManagementCredentials)}

	// Register the user management handler from endpoint using the created dial options.
	err = model_usermanagement.RegisterUserManagementHandlerFromEndpoint(ctx, mux, "localhost:10002", userManagementDialOptions)
	if err != nil {
		log.Fatalf("Error in registering end point. %v", err)
		return nil, err
	}

	return mux, nil
}
