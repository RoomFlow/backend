package services

// Load required packages
import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	model "github.com/RoomFlow/backend/internal/proto/search"
	"github.com/RoomFlow/backend/services/search/helpers"
)

// FilterSearch takes the inputted filters and queries firestore. Returns array of rooms based on filters
func FilterSearch(ctx context.Context, req *model.FilterSearchRequest, firestoreClient *firestore.Client) (*model.FilterSearchResponse, error) {
	// Collection reference
	collRef := firestoreClient.Collection("rooms")

	// Build the query based on the inputted filter
	query := helpers.BuildQuery(collRef, req.Filter)

	log.Println(query)

	// Initialize query iterator
	iter := query.Documents(context.TODO())

	// Return variable
	var rooms []*model.Room

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
			Capacity:   data["Capacity"].(int64),
			RoomType:   model.RoomType(model.RoomType_value[data["RoomType"].(string)]),
			Windows:    Windows,
			Wheelchair: Wheelchair,
			Photos:     Photos,
		}

		// Append document to resulting array
		rooms = append(rooms, &room)
	}

	// Build response
	res := &model.FilterSearchResponse{
		Rooms: rooms,
	}

	return res, nil
}
