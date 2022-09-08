package inspection

type TotalContainer struct {
	InspectedContainers  int `json:"inspectedcontainers,omitempty" validate:"required"`
	ContainersSpotlights int `json:"containersspotlights,omitempty" validate:"required"`
	TreatedContainers    int `json:"treatedcontainers,omitempty" validate:"required"`
	DestroyedContainers  int `json:"destroyedcontainers,omitempty" validate:"required"`
}
