package main

import (
	"log"
	"net/http"

	"github.com/RoomFlow/backend/pkg/config"
	"github.com/RoomFlow/backend/services/apigateway/server"
	"golang.org/x/net/context"
)

func main() {
	// Create a new gateway server.
	gateway, err := server.CreateGateway(context.Background())
	if err != nil {
		log.Fatalf("Error in creating the gateways from other microservice protos : %v", err)
	}

	// Create a new ServeMux.
	serveMux := http.NewServeMux()

	// Registers the handler for the given pattern.
	serveMux.Handle("/", gateway)

	log.Printf("Apigateway deployed on port %s\n", config.ApigatewayPort)

	// listens on the TCP network address and then calls
	// Serve with handler to handle requests on incoming HTTPS connections.
	err = http.ListenAndServeTLS(config.ApigatewayPort, config.SSLCertPath(config.ApigatewayName), config.SSLKeyPath(config.ApigatewayName), serveMux)
	if err != nil {
		log.Fatalf("Error creating an HTTPS connection : %v", err)
	}
}
