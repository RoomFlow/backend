package config

import (
	"log"
	"os"
)

// SpreadsheetID is the id of the room google sheet
var SpreadsheetID string = getEnv("SPREADSHEET_ID", "")

// SearchServicePort is the port that the search service is hosted on
var SearchServicePort string = ":10001"

// UserManagementServicePort is the port that the user management service is hosted on
var UserManagementServicePort string = ":10002"

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	log.Printf("getEnv: returning fallback environment variable for %s:\"%s\"\n", key, fallback)

	return fallback
}
