package repository

import (
	"context"

	"github.com/drizzleent/patients/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetListPatients(context.Context) (*[]model.Patient, error)
	NewPatient(context.Context, *model.Patient) (uuid.UUID, error)
	EditPatient(context.Context, string, *model.Patient) error
	DelPatient(context.Context, string) error
	Load() error
}
