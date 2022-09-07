package inspection

type HomeCondition struct {
	InspectedHome     int `json:"inspected_home,omitempty" validate:"required"`
	ReluctantDwelling int `json:"reluctant_dwelling,omitempty" validate:"required"`
	ClosedHouse       int `json:"closed_home,omitempty" validate:"required"`
	UninhabitedHouse  int `json:"uninhabited_house,omitempty" validate:"required"`
	HousingSpotlights int `json:"housing_spotlights,omitempty" validate:"required"`
	TreatedHousing    int `json:"treated_housing,omitempty" validate:"required"`
}
