package service

import (
	"context"

	"github.com/drizzleent/patients/internal/model"
	"github.com/google/uuid"
)

type ApiService interface {
	GetListPatients(context.Context) (*[]model.ReqPatient, error)
	NewPatient(context.Context, *model.Patient) (uuid.UUID, error)
	EditPatient(context.Context, string, *model.Patient) error
	DelPatient(context.Context, string) error
}
