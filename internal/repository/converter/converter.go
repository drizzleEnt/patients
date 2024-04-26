package converter

import (
	"encoding/json"

	"github.com/drizzleent/patients/internal/model"
	datamodel "github.com/drizzleent/patients/internal/repository/data_model"
)

func FromInmemmroyToModelList(data *[]datamodel.Patient) *[]model.Patient {
	patients := make([]model.Patient, len(*data))
	var patient model.Patient
	for i, v := range *data {
		patient.Fullname = v.Fullname
		patient.Birthday = v.Birthday
		patient.Gender = v.Gender
		patient.Guid = v.Guid
		patients[i] = patient
	}

	return &patients
}

func FromJsonToDataModel(data []byte) (*datamodel.Patient, error) {
	var p datamodel.Patient
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func FromRawJsonToDataModel(data json.RawMessage) (*datamodel.Patient, error) {
	var p datamodel.Patient
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
