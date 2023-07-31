package models

import (
	"gorm.io/gorm"
	"time"
)

type HomeInspection struct {
	gorm.Model
	ID                int64          `json:"id"`
	AddressID         int64          `json:"addressId"`
	Address           Address        `json:"address" validate:"required"`
	Comment           string         `json:"comment" validate:"required"`
	Datetime          time.Time      `json:"datetime" validate:"required"`
	Dni               string         `json:"dni" validate:"required"`
	Latitude          float64        `json:"latitude" validate:"required"`
	Longitude         float64        `json:"longitude" validate:"required"`
	PhotoUrl          string         `json:"photoUrl" validate:"required"`
	NumberInhabitants int32          `json:"numberInhabitants" validate:"required"`
	TypeContainerID   int64          `json:"typeContainerId"`
	TypeContainer     TypeContainer  `json:"typeContainer" validate:"required"`
	HomeConditionID   int64          `json:"homeConditionId"`
	HomeCondition     HomeCondition  `json:"homeCondition" validate:"required"`
	TotalContainerID  int64          `json:"totalContainerId"`
	TotalContainer    TotalContainer `json:"totalContainer" validate:"required"`
	AegyptiFocusID    int64          `json:"aegyptiFocusId"`
	AegyptiFocus      AegyptiFocus   `json:"aegyptiFocus" validate:"required"`
	Larvicide         float32        `json:"larvicide" validate:"required"`
}

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

type Container struct {
	gorm.Model
	ID int64 `json:"id"`
	I  int32 `json:"i" validate:"required"`
	P  int32 `json:"p" validate:"required"`
	T  int32 `json:"t" validate:"required"`
}

type TypeContainer struct {
	gorm.Model
	ID               int64     `json:"id"`
	ElevatedTankID   int64     `json:"elevatedTankId"`
	ElevatedTank     Container `json:"elevatedTank" validate:"required"`
	LowTankID        int64     `json:"lowTankId"`
	LowTank          Container `json:"lowTank" validate:"required"`
	CylinderBarrelID int64     `json:"cylinderBarrelId"`
	CylinderBarrel   Container `json:"cylinderBarrel" validate:"required"`
	BucketTubID      int64     `json:"bucketTubId"`
	BucketTub        Container `json:"bucketTub" validate:"required"`
	TireID           int64     `json:"tireId"`
	Tire             Container `json:"tire" validate:"required"`
	FlowerID         int64     `json:"flowerId"`
	Flower           Container `json:"flower" validate:"required"`
	UselessID        int64     `json:"uselessId"`
	Useless          Container `json:"useless" validate:"required"`
	OthersID         int64     `json:"othersId"`
	Others           Container `json:"others" validate:"required"`
}

type HomeCondition struct {
	gorm.Model
	ID                int64 `json:"id"`
	InspectedHome     int32 `json:"inspectedHome" validate:"required"`
	ReluctantDwelling int32 `json:"reluctantDwelling" validate:"required"`
	ClosedHouse       int32 `json:"closedHouse" validate:"required"`
	UninhabitedHouse  int32 `json:"uninhabitedHouse" validate:"required"`
	HousingSpotlights int32 `json:"housingSpotlights" validate:"required"`
	TreatedHousing    int32 `json:"treatedHousing" validate:"required"`
}

type TotalContainer struct {
	gorm.Model
	ID                   int64 `json:"id"`
	InspectedContainers  int32 `json:"inspectedContainers" validate:"required"`
	ContainersSpotlights int32 `json:"containersSpotlights" validate:"required"`
	TreatedContainers    int32 `json:"treatedContainers" validate:"required"`
	DestroyedContainers  int32 `json:"destroyedContainers" validate:"required"`
}

type AegyptiFocus struct {
	gorm.Model
	ID     int64 `json:"id"`
	Larvae int   `json:"larvae" validate:"required"`
	Pupae  int   `json:"pupae" validate:"required"`
	Adult  int   `json:"adult" validate:"required"`
}
