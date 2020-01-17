package server


import (
	"net/http"
	"log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	userManagement "github.com/RoomFlow/backend/proto/usermanagement"
)

// CreateGateway creates the gateway to access other microservices.
func CreateGateway(ctx context.Context, muxOptions ...runtime.ServeMuxOption) (http.Handler, error) {
    mux := runtime.NewServeMux(muxOptions...)
	
	// Create user management credentials.
	userManagementCredentials, err := credentials.NewClientTLSFromFile("../../../certs/app.crt", "")
	if err != nil {
        log.Fatalf("Error in creating server credentials. %v", err)
        return nil, err
	}
	// Create user management dial options using the created credentials.
	userManagementDialOptions := []grpc.DialOption{grpc.WithTransportCredentials(userManagementCredentials)}

	// Register the user management handler from endpoint using the created dial options.
    err = userManagement.RegisterUserManagementHandlerFromEndpoint(ctx, mux, "localhost:8080", userManagementDialOptions)
    if err != nil {
        log.Fatalf("Error in registering end point. %v", err)
        return nil, err
	}

	return mux, nil
}