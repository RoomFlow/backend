package models

// Room represents one room TODO: Use protobuf struct
type Room struct {
	ID         string   `firestore:"ID,omitempty"`
	RoomType   string   `firestore:"RoomType,omitempty"`
	Capacity   int      `firestore:"Capacity,omitempty"`
	Wheelchair bool     `firestore:"Wheelchair,omitempty"`
	Photos     []string `firestore:"Photos,omitempty"`
	Windows    bool     `firestore:"Windows,omitempty"`
}
