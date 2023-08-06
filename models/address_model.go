package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	ID               int64  `json:"id"`
	FormattedAddress string `json:"formattedAddress" validate:"required"`
	PostalCode       string `json:"postalCode" validate:"required"`
	Country          string `json:"country" validate:"required"`
	Department       string `json:"department" validate:"required"`
	Province         string `json:"province" validate:"required"`
	District         string `json:"district" validate:"required"`
	Urbanization     string `json:"urbanization" validate:"required"`
	Street           string `json:"street" validate:"required"`
	Block            string `json:"block" validate:"required"`
	Lot              string `json:"lot" validate:"required"`
	StreetNumber     string `json:"streetNumber" validate:"required"`
}
