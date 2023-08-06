package models

import (
	"gorm.io/gorm"
	"time"
)

type VectorRecordSummarized struct {
	ID        int64     `json:"id"`
	Latitude  float64   `json:"latitude" validate:"required"`
	Longitude float64   `json:"longitude" validate:"required"`
	Datetime  time.Time `json:"datetime" validate:"required"`
	PhotoUrl  string    `json:"photoUrl" validate:"required"`
}

type VectorRecord struct {
	gorm.Model
	ID        int64     `json:"id"`
	AddressID int64     `json:"addressId"`
	Address   Address   `json:"address" validate:"required"`
	Comment   string    `json:"comment" validate:"required"`
	Datetime  time.Time `json:"datetime" validate:"required"`
	Latitude  float64   `json:"latitude" validate:"required"`
	Longitude float64   `json:"longitude" validate:"required"`
	PhotoUrl  string    `json:"photoUrl" validate:"required"`
}
