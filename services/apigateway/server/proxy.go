package server

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/RoomFlow/backend/pkg/config"
	model_search "github.com/RoomFlow/backend/pkg/proto/search"
	model_usermanagement "github.com/RoomFlow/backend/pkg/proto/usermanagement"
)

// CreateGateway creates the gateway to access other microservices.
func CreateGateway(ctx context.Context, muxOptions ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(muxOptions...)

	var err error

	// Create client credentials with certificate if we are not on development
	var searchCredentials credentials.TransportCredentials
	if config.Environment == "development" {
		searchCredentials = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	} else {
		searchCredentials, err = credentials.NewClientTLSFromFile(config.SSLCertPath(config.SearchServiceName), "")
		if err != nil {
			log.Fatalf("Error in creating server credentials. %v", err)
			return nil, err
		}
	}

	// Create search service dial options using the created credentials.
	searchDialOptions := []grpc.DialOption{grpc.WithTransportCredentials(searchCredentials)}

	// Register the search service handler from endpoint using the created dial options.
	err = model_search.RegisterSearchHandlerFromEndpoint(ctx, mux, config.SearchEndpoint, searchDialOptions)
	if err != nil {
		log.Fatalf("Error in registering end point. %v", err)
		return nil, err
	}

	// Create client credentials with certificate if we are not on development
	var usermanagementCredentials credentials.TransportCredentials
	if config.Environment == "development" {
		usermanagementCredentials = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	} else {
		usermanagementCredentials, err = credentials.NewClientTLSFromFile(config.SSLCertPath(config.UsermanagementServiceName), "")
		if err != nil {
			log.Fatalf("Error in creating server credentials. %v", err)
			return nil, err
		}
	}

	// Create usermanagement service dial options using the created credentials.
	userManagementDialOptions := []grpc.DialOption{grpc.WithTransportCredentials(usermanagementCredentials)}

	// Register the usermanagement service handler from endpoint using the created dial options.
	err = model_usermanagement.RegisterUserManagementHandlerFromEndpoint(ctx, mux, config.UsermanagementEndpoint, userManagementDialOptions)
	if err != nil {
		log.Fatalf("Error in registering end point. %v", err)
		return nil, err
	}

	return mux, nil
}
