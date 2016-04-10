package main

// Represents a character.
type Character struct {
	ID string `bson:"_id" json:"_id"`
	ProID string `bson:"pro_id" json:"pro_id"`
	PlayerNum int `bson:"num" json:"num"`
	Experience int `bson:"experience" json:"experience"`
	Gold int `bson:"gold" json:"gold"`
}
