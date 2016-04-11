package models

// For auto-incrementing player numbers.
type Counter struct {
	ID  string `bson:"_id" json:"id"`
	Seq int    `bson:"count" json:"count"`
}
