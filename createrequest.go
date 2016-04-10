package main

// Represents a character creation request.
type CreateRequest struct {
	Data string `bson:"data" json:"data"`
}
