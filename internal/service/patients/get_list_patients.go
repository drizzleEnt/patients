package patients

import (
	"context"

	"github.com/drizzleent/patients/internal/model"
)

func (s *srv) GetListPatients(ctx context.Context) (*[]model.Patient, error) {

	res, err := s.repo.GetListPatients(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
