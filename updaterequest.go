package main

type UpdateRequest struct {
	ID         string `bson:"_id" json:"_id"`
	ProID      string `bson:"pro_id" json:"pro_id"`
	Experience int    `bson:"experience" json:"experience"`
	Gold       int    `bson:"gold" json:"gold"`
}
