package patients

import (
	"context"

	"github.com/drizzleent/patients/internal/converter"
	"github.com/drizzleent/patients/internal/model"
)

func (s *srv) GetListPatients(ctx context.Context) (*[]model.ReqPatient, error) {

	patients, err := s.repo.GetListPatients(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]model.ReqPatient, 0, len(*patients))
	for _, v := range *patients {
		res = append(res, *converter.FromPatientToReq(&v))
	}
	return &res, nil
}
