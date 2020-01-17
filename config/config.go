package config

import (
	"log"
	"os"
)

// SpreadsheetID is the id of the room google sheet
var SpreadsheetID string = getEnv("SPREADSHEET_ID", "")

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	log.Printf("getEnv: returning fallback environment variable, %v\n", fallback)

	return fallback
}
