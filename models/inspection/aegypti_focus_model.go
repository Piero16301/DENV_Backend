package inspection

type AegyptiFocus struct {
	Larvae int `json:"larvae,omitempty" validate:"required"`
	Pupae  int `json:"pupae,omitempty" validate:"required"`
	Adult  int `json:"adult,omitempty" validate:"required"`
}
