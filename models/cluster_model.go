package models

type Cluster struct {
	Id     int            `json:"id,omitempty" validate:"required"`
	Points []ClusterPoint `json:"points,omitempty" validate:"required"`
}
