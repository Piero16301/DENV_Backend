package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type VectorRecord struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Address   Address            `json:"address,omitempty" validate:"required"`
	Comment   string             `json:"comment,omitempty" validate:"required"`
	Datetime  time.Time          `json:"datetime,omitempty" validate:"required"`
	Latitude  float32            `json:"latitude,omitempty" validate:"required"`
	Longitude float32            `json:"longitude,omitempty" validate:"required"`
	PhotoURL  string             `json:"photourl,omitempty" validate:"required"`
}
