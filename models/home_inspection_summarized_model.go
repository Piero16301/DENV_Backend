package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HomeInspectionSummarized struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Latitude  float32            `json:"latitude,omitempty" validate:"required"`
	Longitude float32            `json:"longitude,omitempty" validate:"required"`
	Datetime  time.Time          `json:"datetime,omitempty" validate:"required"`
	PhotoURL  string             `json:"photourl,omitempty" validate:"required"`
}
