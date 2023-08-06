package models

import "time"

type ClusterPoint struct {
	ID        int64     `json:"id" validate:"required"`
	Latitude  float64   `json:"latitude" validate:"required"`
	Longitude float64   `json:"longitude" validate:"required"`
	Datetime  time.Time `json:"datetime" validate:"required"`
}
