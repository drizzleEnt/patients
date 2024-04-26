package patients

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/drizzleent/patients/internal/model"
)

const (
	filePath = "data/list_patients.json"
)

type repo struct {
	data    []byte
	allJson []json.RawMessage
}

func NewRepository() *repo {
	return &repo{}
}

func (r *repo) Load() error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fd, err := file.Stat()
	if err != nil {
		return err
	}
	r.data = make([]byte, fd.Size())
	_, err = file.Read(r.data)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	allJson := []json.RawMessage{}

	err = json.Unmarshal(r.data, &allJson)
	if err != nil {
		return err
	}

	r.allJson = allJson
	return nil
}

func (r *repo) GetListPatients(ctx context.Context) (*[]model.Patient, error) {
	res := make([]model.Patient, len(r.allJson))

	for i, v := range r.allJson {
		err := json.Unmarshal(v, &res[i])

		if err != nil {
			return nil, err
		}

	}

	return &res, nil
}

func (r *repo) NewPatient(ctx context.Context) {

}

func (r *repo) EditPatient(ctx context.Context) {

}

func (r *repo) DelPatient(ctx context.Context) {

}
