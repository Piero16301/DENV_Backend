package inspection

import "DENV_Backend/models/inspection/containers"

type TypeContainers struct {
	ElevatedTank   containers.ElevatedTank   `json:"elevatedTank,omitempty" validate:"required"`
	LowTank        containers.LowTank        `json:"lowTank,omitempty" validate:"required"`
	CylinderBarrel containers.CylinderBarrel `json:"cylinderBarrel,omitempty" validate:"required"`
	BucketTub      containers.BucketTub      `json:"bucketTub,omitempty" validate:"required"`
	Tire           containers.Tire           `json:"tire,omitempty" validate:"required"`
	Flower         containers.Flower         `json:"flower,omitempty" validate:"required"`
	Useless        containers.Useless        `json:"useless,omitempty" validate:"required"`
	Others         containers.Others         `json:"others,omitempty" validate:"required"`
}
