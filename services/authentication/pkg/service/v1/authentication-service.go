package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/RoomFlow/backend/services/authentication/pkg/api/v1"
)

// toDoServiceServer is implementation of v1.AuthenticationServer proto interface
type authenticationServiceServer struct {
	// empty.
}

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// NewAuthenticationServiceServer creates Authentication service.
func NewAuthenticationServiceServer() v1.AuthenticationServiceServer {
	return &authenticationServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *authenticationServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// func main() {
// 	listener, err := net.Listen("tcp", ":4040")
// 	if err != nil {
// 		panic(err)
// 	}

// 	srv := grpc.NewServer()
// 	proto.RegisterLoginServiceServer(srv, &server{})
// 	reflection.Register(srv)

// 	if e := srv.Serve(listener); e != nil {
// 		panic(e)
// 	}

// }

// Create a login request.
func (s *authenticationServiceServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// TODO: call firebase function for authentication
	// TODO: return server code from firebase function for authentication

	return &v1.LoginResponse{
		Api: apiVersion, 
		Status: "400",
	}, nil
}
