package patients

import (
	"context"
)

func (s *srv) DelPatient(ctx context.Context, id string) error {
	err := s.repo.DelPatient(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
