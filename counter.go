package main

type Counter struct {
	ID string `bson:"_id", json:"id"`
	Seq int `bson:"count", json:"count"`
}
