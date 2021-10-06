package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitave"
)

type Order struct {
	ID         primitave.ObjectID `bson:"_id"`
	Order_Date time.Time          `json"order_date" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Order_id   string             `json:"order_id" validate:"required"`
	Table_id   *string            `json:"table_id" validate:"requried"`
}