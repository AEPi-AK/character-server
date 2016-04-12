package models

// Represents a character.
type Character struct {
	ID         string `bson:"_id" json:"id"`
	ProID      string `bson:"pro_id" json:"pro_id"`
	PlayerNum  int    `bson:"num" json:"num"`
	Experience int    `bson:"experience" json:"experience"`
	Race string `bson:"race" json:"race"`
	Strength int `bson:"strength" json:"strength"`
	Dexterity int `bson:"dexterity" json:"dexterity"`
	Wisdom int `bson:"wisdom" json:"wisdom"`
}
