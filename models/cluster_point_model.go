package models

import "time"

type ClusterPoint struct {
	Id        string    `json:"id,omitempty" validate:"required"`
	Latitude  float32   `json:"latitude,omitempty" validate:"required"`
	Longitude float32   `json:"longitude,omitempty" validate:"required"`
	Datetime  time.Time `json:"datetime,omitempty" validate:"required"`
}
