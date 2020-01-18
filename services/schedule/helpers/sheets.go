package helpers

import (
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	config "github.com/RoomFlow/backend/internal/config"
	models "github.com/RoomFlow/backend/internal/models"
)

// FetchRoomData fetches room data from McMaster room directory
func FetchRoomData() []models.Room {
	log.Println("Creating new Sheets service")

	// Start new google sheets api service
	sheetsService, err := sheets.NewService(context.TODO(), option.WithScopes(sheets.SpreadsheetsReadonlyScope))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Get spreadsheet id from config
	spreadsheetID := config.SpreadsheetID

	// Range of columns to fetch
	readRange := "ClassroomData!A2:EZ"

	log.Println("Fetching spreadsheet data")

	// Fetch data
	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Return array
	var rooms []models.Room

	// Iterate through returned rows
	for _, row := range resp.Values {
		// Some entries have "/" which breaks the addition of the document in firestore
		ID := strings.ReplaceAll(row[0].(string), "/", "-")

		// Convert capacity string to int
		Capacity, err := strconv.Atoi(row[7].(string))
		if err != nil {
			log.Printf("Unable to convert capacity string to int: %s\n", err)
		}

		RoomType := strings.ToUpper(strings.ReplaceAll(row[6].(string), " ", "_"))

		Wheelchair := false
		if len(row[16].(string)) > 0 {
			Wheelchair = true
		}

		var Photos []string

		rowLength := len(row)

		if rowLength > 153 {
			Photos = append(Photos, row[153].(string))
		}

		if rowLength > 154 {
			Photos = append(Photos, row[154].(string))
		}

		if rowLength > 155 {
			Photos = append(Photos, row[155].(string))
		}

		Windows := false
		if len(row[151].(string)) > 0 {
			Windows = true
		}

		// Create room struct
		roomData := models.Room{
			ID:         ID,
			RoomType:   RoomType,
			Capacity:   Capacity,
			Wheelchair: Wheelchair,
			Photos:     Photos,
			Windows:    Windows,
		}

		// Add room data to array
		rooms = append(rooms, roomData)
	}

	return rooms
}
