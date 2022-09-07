package containers

type LowTank struct {
	I int `json:"i,omitempty" validate:"required"`
	P int `json:"p,omitempty" validate:"required"`
	T int `json:"t,omitempty" validate:"required"`
}
