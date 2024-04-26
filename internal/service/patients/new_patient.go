package patients

import (
	"context"

	"github.com/drizzleent/patients/internal/model"
	"github.com/google/uuid"
)

func (s *srv) NewPatient(ctx context.Context, p *model.Patient) (uuid.UUID, error) {
	id, err := s.repo.NewPatient(ctx, p)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
