package helpers

import (
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/RoomFlow/backend/config"
)

// Room represents one room
type Room struct {
	ID       string `firestore:"id,omitempty"`
	RoomType string `firestore:"roomType,omitempty"`
	Capacity int    `firestore:"capacity,omitempty"`
}

// FetchRoomData fetches room data from McMaster room directory
func FetchRoomData() []Room {
	log.Println("Creating new Sheets service")

	// Start new google sheets api service
	sheetsService, err := sheets.NewService(context.TODO(), option.WithScopes(sheets.SpreadsheetsReadonlyScope))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Get spreadsheet id from config
	spreadsheetID := config.SpreadsheetID

	// Range of columns to fetch
	readRange := "ClassroomData!A2:H"

	log.Println("Fetching spreadsheet data")

	// Fetch data
	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Return array
	var rooms []Room

	// Iterate through returned rows
	for _, row := range resp.Values {
		// Some entries have "/" which breaks the addition of the document in firestore
		ID := strings.ReplaceAll(row[0].(string), "/", "-")

		// Convert capacity string to int
		Capacity, err := strconv.Atoi(row[7].(string))
		if err != nil {
			log.Printf("Unable to concert capacity string to int: %s\n", err)
		}

		// Create room struct
		roomData := Room{
			ID:       ID,
			RoomType: row[6].(string),
			Capacity: Capacity,
		}

		// Add room data to array
		rooms = append(rooms, roomData)
	}

	return rooms
}
