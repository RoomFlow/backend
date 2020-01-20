package config

import (
	"log"
	"os"
)

// SpreadsheetID is the id of the room google sheet
var SpreadsheetID string = getEnv("SPREADSHEET_ID", "")

// ApigatewayName is the name of apigateway
var ApigatewayName string = "apigateway"

// ApigatewayPort is the port that apigateway is deployed on (NAMED "PORT" BECAUSE HEROKU)
var ApigatewayPort string = getEnv("APIGATEWAY_PORT", ":443")

//SearchServiceName is the name of the search service
var SearchServiceName string = "search"

// SearchServicePort is the port that the search service is deployed on
var SearchServicePort string = getEnv("SEARCH_PORT", ":10001")

// SearchEndpoint is the endpoint which the search service can be called
var SearchEndpoint string = getEnv("SEARCH_ENDPOINT", "localhost"+SearchServicePort)

// UsermanagementServiceName is the name of the usermanagement service
var UsermanagementServiceName string = "usermanagement"

// UsermanagementServicePort is the port that the usermanagement service is deployed on
var UsermanagementServicePort string = getEnv("USERMANAGEMENT_PORT", ":10002")

// UsermanagementEndpoint is the endpoint which the usermanagement service can be called
var UsermanagementEndpoint string = getEnv("USERMANAGEMENT_ENDPOINT", "localhost"+UsermanagementServicePort)

// Environment is use to determine whether we are in production or development mode
var Environment string = getEnv("GO_ENV", "development")

// SSLCertPath is the path where the ssl cert is located
func SSLCertPath(serviceName string) string {
	// Returns default cert if env var not set
	return getEnv("SSL_CERT_PATH", "internal/certs/"+serviceName+"/app.crt")
}

// SSLKeyPath is the path where the ssl cert is located
func SSLKeyPath(serviceName string) string {
	// Returns default cert if env var not set
	return getEnv("SSL_KEY_PATH", "internal/certs/"+serviceName+"/app.key")
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	log.Printf("getEnv: returning fallback environment variable for %s:\"%s\"\n", key, fallback)

	return fallback
}
