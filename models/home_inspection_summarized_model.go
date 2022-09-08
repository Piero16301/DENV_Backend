package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HomeInspectionSummarized struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Latitude  float32            `json:"latitude,omitempty" validate:"required"`
	Longitude float32            `json:"longitude,omitempty" validate:"required"`
}
