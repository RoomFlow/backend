package helpers

import (
	"errors"
	"log"
	"reflect"

	"cloud.google.com/go/firestore"

	model "github.com/RoomFlow/backend/internal/proto/search"
)

// BuildQuery builds the query based on the filter
func BuildQuery(collection *firestore.CollectionRef, filter *model.Filter) (firestore.Query, error) {
	log.Printf("Building query with filter %v\n", filter)

	// Reflect value
	filterValue := reflect.ValueOf(filter).Elem()

	// Initialize firestore query to append to
	var query firestore.Query

	if !filterValue.IsValid() {
		err := errors.New("Filter is not valid")
		log.Println(err)
		return query, err
	}

	// Reflect value type
	typeOfFilter := filterValue.Type()

	// This is so we can know when to initialize the query based on the collection ref
	queryInitialized := false

	// Iterate through all fields in filter struct
	for i := 0; i < filterValue.NumField()-3; i++ {
		// Reflect field object
		field := filterValue.Field(i)
		// Field value interface
		fieldValue := field.Interface()

		// Check if filter value is not nil
		if !field.IsZero() {
			// Name of field (string)
			name := typeOfFilter.Field(i).Name
			// Default comparator is "==" since we can have only one inequality comparison in the query
			comparison := comparisons["EQUAL_TO"]

			// Check if filter field is "Capacity" and apply the comparison symbol which was passed in the filter
			if name == "Capacity" {
				comparison = comparisons[filter.GetCapacity().GetComparison().String()]
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

	if queryInitialized {
		return query, nil
	}

	err := errors.New("Filter is not valid")
	log.Println(err)
	return query, err
}
