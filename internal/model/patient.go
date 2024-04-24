package model

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	Fullname string
	Birthday time.Time
	Gender   int
	Guid     uuid.UUID
}
