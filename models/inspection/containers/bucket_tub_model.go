package containers

type BucketTub struct {
	I int `json:"i,omitempty" validate:"required"`
	P int `json:"p,omitempty" validate:"required"`
	T int `json:"t,omitempty" validate:"required"`
}
