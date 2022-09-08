package models

type Address struct {
	FormattedAddress string `json:"formattedaddress,omitempty" validate:"required"`
	PostalCode       string `json:"postalcode,omitempty" validate:"required"`
	Country          string `json:"country,omitempty" validate:"required"`
	Department       string `json:"department,omitempty" validate:"required"`
	Province         string `json:"province,omitempty" validate:"required"`
	District         string `json:"district,omitempty" validate:"required"`
	Urbanization     string `json:"urbanization,omitempty" validate:"required"`
	Street           string `json:"street,omitempty" validate:"required"`
	Block            string `json:"block,omitempty" validate:"required"`
	Lot              int    `json:"lot,omitempty" validate:"required"`
	StreetNumber     int    `json:"streetnumber,omitempty" validate:"required"`
}