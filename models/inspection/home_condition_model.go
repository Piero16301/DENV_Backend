package inspection

type HomeCondition struct {
	InspectedHome     int `json:"inspectedhome,omitempty" validate:"required"`
	ReluctantDwelling int `json:"reluctantdwelling,omitempty" validate:"required"`
	ClosedHouse       int `json:"closedhome,omitempty" validate:"required"`
	UninhabitedHouse  int `json:"uninhabitedhouse,omitempty" validate:"required"`
	HousingSpotlights int `json:"housingspotlights,omitempty" validate:"required"`
	TreatedHousing    int `json:"treatedhousing,omitempty" validate:"required"`
}
