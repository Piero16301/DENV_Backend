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
	PhotoURL  string             `json:"photourl,omitempty" validate:"required"`

	NumberInhabitants int                       `json:"numberinhabitants,omitempty" validate:"required"`
	HomeCondition     inspection.HomeCondition  `json:"homecondition,omitempty" validate:"required"`
	TypeContainers    inspection.TypeContainers `json:"typecontainers,omitempty" validate:"required"`
	TotalContainer    inspection.TotalContainer `json:"totalcontainer,omitempty" validate:"required"`
	AegyptiFocus      inspection.AegyptiFocus   `json:"aegyptifocus,omitempty" validate:"required"`
}
