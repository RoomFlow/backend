package services

// Load required packages
import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	model "github.com/RoomFlow/backend/pkg/proto/search"
	"github.com/RoomFlow/backend/services/search/helpers"
)

// Filter takes the inputted filters and queries firestore. Returns array of rooms based on filters
func Filter(ctx context.Context, req *model.FilterRequest, firestoreClient *firestore.Client) (*model.FilterResponse, error) {
	// Collection reference
	collRef := firestoreClient.Collection("rooms")

	// Build the query based on the inputted filter
	query, err := helpers.BuildQuery(collRef, req.Filter)
	if err != nil {
		return nil, err
	}

	log.Printf("Built query: %v\n", query)

	// Initialize query iterator
	iter := query.Documents(context.TODO())

	// Return variable
	var rooms []*model.Room

	log.Println("Iterating through query results")

	// Loop through iterator to get all resulting documents
	for {
		// Get next document in iterator
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		// Document Map
		data := doc.Data()

		var Windows bool
		if val, ok := data["Windows"]; ok {
			Windows = val.(bool)
		}

		var Wheelchair bool
		if val, ok := data["Wheelchair"]; ok {
			Wheelchair = val.(bool)
		}

		var Photos []string
		if val, ok := data["Photos"]; ok {
			s := make([]string, len(data["Photos"].([]interface{})))
			for i, v := range val.([]interface{}) {
				s[i] = fmt.Sprint(v)
			}
			Photos = s
		}

		// TODO: Convert document to Room protobuf model automatically
		room := model.Room{
			ID:         data["ID"].(string),
			Building:   model.Building(model.Building_value[data["Building"].(string)]),
			RoomNumber: data["RoomNumber"].(string),
			Capacity:   data["Capacity"].(int64),
			RoomType:   model.RoomType(model.RoomType_value[data["RoomType"].(string)]),
			Windows:    Windows,
			Wheelchair: Wheelchair,
			Photos:     Photos,
		}

		// Append document to resulting array
		rooms = append(rooms, &room)
	}

	log.Println("Done searching, returning response")

	// Build response
	res := &model.FilterResponse{
		Rooms: rooms,
	}

	return res, nil
}
