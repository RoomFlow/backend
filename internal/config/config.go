package config

import (
	"log"
	"os"
)

// SpreadsheetID is the id of the room google sheet
var SpreadsheetID string = getEnv("SPREADSHEET_ID", "")

// ApigatewayPort is the port that apigateway is deployed on
var ApigatewayPort string = getEnv("APIGATEWAY_PORT", ":443")

// SearchServicePort is the port that the search service is deployed on
var SearchServicePort string = getEnv("SEARCH_PORT", ":10001")

// SearchEndpoint is the endpoint which the search service can be called
var SearchEndpoint string = getEnv("SEARCH_ENDPOINT", "localhost"+SearchServicePort)

// UsermanagementServicePort is the port that the usermanagement service is deployed on
var UsermanagementServicePort string = getEnv("USERMANAGEMENT_PORT", ":10002")

// UsermanagementEndpoint is the endpoint which the usermanagement service can be called
var UsermanagementEndpoint string = getEnv("USERMANAGEMENT_ENDPOINT", "localhost"+UsermanagementServicePort)

// SSLCertPath is the path where the ssl cert is located
var SSLCertPath string = getEnv("SSL_CERT_PATH", "internal/certs/app.crt")

// SSLKeyPath is the path where the ssl key is located
var SSLKeyPath string = getEnv("SSL_Key_PATH", "internal/certs/app.key")

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	log.Printf("getEnv: returning fallback environment variable for %s:\"%s\"\n", key, fallback)

	return fallback
}
