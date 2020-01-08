package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	userManagement "github.com/RoomFlow/backend/proto/usermanagement"
)


func main() {
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := userManagement.NewUserManagementClient(conn)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	loginResponse, loginErr := c.LoginUser(ctx, &userManagement.LoginRequest{
		Username: "username",
		Password: "password",
	})
	registerResponse, registerErr := c.RegisterUser(ctx, &userManagement.RegisterRequest{
		Username: "username",
		Password: "password",
	})

	if loginErr != nil {
		log.Fatalf("could not login: %v", loginErr)
		return
	} else {
		log.Printf("Login token: %s", loginResponse.GetToken())
	}

	if registerErr != nil {
		log.Fatalf("could not register: %v", registerErr)
		return
	} else {
		log.Printf("Register token: %s", registerResponse.GetToken())
	}
}