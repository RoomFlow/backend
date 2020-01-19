package controllers

// Load required packages
import (
	"context"

	"cloud.google.com/go/firestore"
	model "github.com/RoomFlow/backend/internal/proto/search"
	"github.com/RoomFlow/backend/services/search/services"
)

// SearchServer is the search server
type SearchServer struct {
	FirestoreClient *firestore.Client
}

// Filter searches based on inputted filter
func (server *SearchServer) Filter(ctx context.Context, in *model.FilterRequest) (*model.FilterResponse, error) {
	return services.Filter(ctx, in, server.FirestoreClient)
}
