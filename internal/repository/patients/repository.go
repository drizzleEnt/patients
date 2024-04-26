package patients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/drizzleent/patients/internal/model"
	"github.com/drizzleent/patients/internal/repository/converter"
	datamodel "github.com/drizzleent/patients/internal/repository/data_model"
	"github.com/google/uuid"
)

const (
	filePath = "data/list_patients.json"
)

type repo struct {
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
	data := make([]byte, fd.Size())
	_, err = file.Read(data)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	allJson := []json.RawMessage{}

	err = json.Unmarshal(data, &allJson)
	if err != nil {
		return err
	}
	r.allJson = allJson
	return nil
}

func (r *repo) GetListPatients(_ context.Context) (*[]model.Patient, error) {
	res := make([]datamodel.Patient, len(r.allJson))

	for i, v := range r.allJson {
		err := json.Unmarshal(v, &res[i])

		if err != nil {
			return nil, err
		}

	}

	return converter.FromInmemmroyToModelList(&res), nil
}

func (r *repo) NewPatient(ctx context.Context, p *model.Patient) (uuid.UUID, error) {
	id := uuid.New()
	p.Guid = id
	data, err := json.Marshal(p)

	if err != nil {
		return uuid.Nil, err
	}

	r.allJson = append(r.allJson, data)
	for _, v := range r.allJson {
		fmt.Println(string(v))
	}

	err = r.saveInFile()
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *repo) EditPatient(ctx context.Context) {

}

func (r *repo) DelPatient(ctx context.Context) {

}

func (r *repo) saveInFile() error {
	data, err := json.Marshal(r.allJson)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, data, fs.ModeAppend)
	if err != nil {
		return err
	}

	return nil

}
