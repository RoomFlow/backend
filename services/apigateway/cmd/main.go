package main

import (
	"github.com/RoomFlow/backend/services/apigateway/server"
	"golang.org/x/net/context"
	"log"
	"net/http"
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

	log.Println("Apigateway deployed on port 443")

	// listens on the TCP network address and then calls
	// Serve with handler to handle requests on incoming HTTPS connections.
	err = http.ListenAndServeTLS(":443", "internal/certs/app.crt", "internal/certs/app.key", serveMux)
	if err != nil {
		log.Fatalf("Error creating an HTTPS connection : %v", err)
	}
}
