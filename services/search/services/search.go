package services

// Load required packages
import (
	"context"
	"fmt"
	"log"
	"reflect"

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
	query := buildQuery(collRef, req.Filter)

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

// buildQuery builds the query based on the filter
func buildQuery(collection *firestore.CollectionRef, filter *model.Filter) firestore.Query {
	// Reflect value
	filterValue := reflect.ValueOf(filter).Elem()
	// Reflect value type
	typeOfFilter := filterValue.Type()

	// Initialize firestore query to append to
	var query firestore.Query

	// This is so we can know when to initialize the query based on the collection ref
	queryInitialized := false

	// Iterate through all fields in filter struct
	for i := 0; i < filterValue.NumField()-3; i++ {
		// Reflect field object
		field := filterValue.Field(i)
		// Field value interface
		fieldValue := field.Interface()
		// Check if filter is not nil
		if !field.IsZero() {
			// Name of field (string)
			name := typeOfFilter.Field(i).Name
			// Default comparator is "==" since we can have only one inequality comparison in the query
			comparison := helpers.Comparisons["EQUAL_TO"]

			// Check if filter field is "Capacity" and apply the comparison symbol which was passed in the filter
			if name == "Capacity" {
				comparison = helpers.Comparisons[filter.GetCapacity().GetComparison().String()]
				fieldValue = filter.GetCapacity().GetSize()
			}

			if name == "RoomType" {
				fieldValue = fieldValue.(model.RoomType).String()
			}

			// Here we check if the query was already initialized for the first time
			if !queryInitialized {
				query = collection.Where(name, comparison, fieldValue)
				queryInitialized = true
			} else {
				// Append where clause to query
				query = query.Where(name, comparison, fieldValue)
			}
		}
	}

	return query
}
