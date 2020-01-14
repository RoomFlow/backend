// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/RoomFlow/backend/services/apigateway/models"
	"github.com/RoomFlow/backend/services/apigateway/restapi/operations"
	"github.com/RoomFlow/backend/services/apigateway/restapi/operations/login"
	"github.com/RoomFlow/backend/services/apigateway/restapi/operations/register"

	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	userManagement "github.com/RoomFlow/backend/proto/usermanagement"
)

//go:generate swagger generate server --target ../../apigateway --name APIGateway --spec ../swagger.yml

func configureFlags(api *operations.APIGatewayAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.APIGatewayAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()


	api.LoginPostLoginHandler = login.PostLoginHandlerFunc(func(params login.PostLoginParams) middleware.Responder {
		// Set up a connection to the server.
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := userManagement.NewUserManagementClient(conn)


		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		loginResponse, loginErr := c.LoginUser(ctx, &userManagement.LoginRequest{
			Username: *params.Body.Username,
			Password: *params.Body.Password,
		})

		if loginErr != nil {
			log.Fatalf("could not login: %v", loginErr)
			return login.NewPostLoginDefault(401)
		} else {
			log.Printf("Login token: %s", loginResponse.GetToken())
			token := loginResponse.GetToken()
			return login.NewPostLoginOK().WithPayload(&models.Token{Accesstoken: &token})
		}
	})
	
	api.RegisterPostRegisterHandler = register.PostRegisterHandlerFunc(func(params register.PostRegisterParams) middleware.Responder {
		// Set up a connection to the server.
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := userManagement.NewUserManagementClient(conn)


		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		registerResponse, registerErr := c.RegisterUser(ctx, &userManagement.RegisterRequest{
			Username: *params.Body.Username,
			Password: *params.Body.Password,
		})

		if registerErr != nil {
			log.Fatalf("could not register: %v", registerErr)
			return register.NewPostRegisterDefault(401)
		} else {
			log.Printf("Register token: %s", registerResponse.GetToken())
			token := registerResponse.GetToken()
			return register.NewPostRegisterOK().WithPayload(&models.Token{Accesstoken: &token})
		}
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
