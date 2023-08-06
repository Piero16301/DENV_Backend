package models

type Cluster struct {
	Id     int            `json:"id" validate:"required"`
	Points []ClusterPoint `json:"points" validate:"required"`
}
