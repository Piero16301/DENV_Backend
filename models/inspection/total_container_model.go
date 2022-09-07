package inspection

type TotalContainer struct {
	InspectedContainers  int `json:"inspected_containers,omitempty" validate:"required"`
	ContainersSpotlights int `json:"containers_spotlights,omitempty" validate:"required"`
	TreatedContainers    int `json:"treated_containers,omitempty" validate:"required"`
	DestroyedContainers  int `json:"destroyed_containers,omitempty" validate:"required"`
}
