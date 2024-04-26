package service

import (
	"context"

	"github.com/drizzleent/patients/internal/model"
)

type ApiService interface {
	GetListPatients(context.Context) (*[]model.Patient, error)
	NewPatient(context.Context)
	EditPatient(context.Context)
	DelPatient(context.Context)
}
