package patients

import (
	"context"

	"github.com/drizzleent/patients/internal/model"
)

func (s *srv) EditPatient(ctx context.Context, id string, p *model.Patient) error {
	err := s.repo.EditPatient(ctx, id, p)
	if err != nil {
		return err
	}

	return nil
}
