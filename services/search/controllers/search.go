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

// FilterSearch searches based on inputted filter
func (server *SearchServer) FilterSearch(ctx context.Context, in *model.FilterSearchRequest) (*model.FilterSearchResponse, error) {
	return services.FilterSearch(ctx, in, server.FirestoreClient)
}
