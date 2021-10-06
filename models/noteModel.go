package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitave"
)

type Note struct {
	ID         primitave.ObjectID `bson:"_id"`
	Text       string             `json:"text"`
	Title      string             `json:"title"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Note_id    string             `json:"note_id"`
}
