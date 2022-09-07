package models

import (
	"DENV_Backend/models/inspection"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HomeInspection struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Address   Address            `json:"address,omitempty" validate:"required"`
	Comment   string             `json:"comment,omitempty" validate:"required"`
	Datetime  time.Time          `json:"datetime,omitempty" validate:"required"`
	DNI       string             `json:"dni,omitempty" validate:"required"`
	Latitude  float32            `json:"latitude,omitempty" validate:"required"`
	Longitude float32            `json:"longitude,omitempty" validate:"required"`
	PhotoURL  string             `json:"photo_url,omitempty" validate:"required"`

	NumberInhabitants int                       `json:"number_inhabitants,omitempty" validate:"required"`
	HomeCondition     inspection.HomeCondition  `json:"home_condition,omitempty" validate:"required"`
	TypeContainers    inspection.TypeContainers `json:"type_containers,omitempty" validate:"required"`
	TotalContainer    inspection.TotalContainer `json:"total_containers,omitempty" validate:"required"`
	AegyptiFocus      inspection.AegyptiFocus   `json:"aegypti_focus,omitempty" validate:"required"`
}
