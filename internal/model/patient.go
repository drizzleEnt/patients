package model

import (
	"github.com/google/uuid"
)

type Patient struct {
	Fullname      string    `json:"fullname"`
	Birthday      string    `json:"birthday"`
	Gender        int       `json:"gender"`
	Guid          uuid.UUID `json:"guid"`
	IsGenderValid bool
}

type ReqPatient struct {
	Fullname string    `json:"fullname"`
	Birthday string    `json:"birthday"`
	Gender   int       `json:"gender"`
	Guid     uuid.UUID `json:"guid"`
}
