package models

import "time"

// Represents a character.
type Character struct {
	ID         string `bson:"_id" json:"id"`
	ProID      string `bson:"pro_id" json:"pro_id"`
	PlayerNum  int    `bson:"number" json:"number"`
	Experience int    `bson:"points" json:"points"`
	Race string `bson:"race" json:"race"`
	Strength int `bson:"strength" json:"strength"`
	Dexterity int `bson:"dexterity" json:"dexterity"`
	Wisdom int `bson:"wisdom" json:"wisdom"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"` 
}
