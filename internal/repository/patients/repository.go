package patients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"sync"

	"github.com/drizzleent/patients/internal/model"
	"github.com/drizzleent/patients/internal/repository/converter"
	datamodel "github.com/drizzleent/patients/internal/repository/data_model"
	"github.com/google/uuid"
)

const (
	filePath = "data/list_patients.json"
)

type repo struct {
	m        sync.RWMutex
	patients map[string]datamodel.Patient
}

func NewRepository() *repo {
	r := &repo{}
	return r
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
	r.patients = map[string]datamodel.Patient{}
	for _, v := range allJson {
		p, err := converter.FromRawJsonToDataModel(v)
		if err != nil {
			return err
		}
		r.patients[p.Guid.String()] = *p
	}
	return nil
}

func (r *repo) GetListPatients(_ context.Context) (*[]model.Patient, error) {
	r.m.Lock()
	defer r.m.Unlock()
	res := make([]datamodel.Patient, 0, len(r.patients))

	for _, v := range r.patients {
		res = append(res, v)
	}

	return converter.FromInmemmroyToModelList(&res), nil
}

func (r *repo) NewPatient(_ context.Context, p *model.Patient) (uuid.UUID, error) {
	id := uuid.New()
	p.Guid = id
	data, err := json.Marshal(p)

	if err != nil {
		return uuid.Nil, err
	}

	dataPatient, err := converter.FromJsonToDataModel(data)
	if err != nil {
		return uuid.Nil, err
	}
	r.m.Lock()

	r.patients[id.String()] = *dataPatient

	err = r.saveInFile()
	r.m.Unlock()

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *repo) EditPatient(ctx context.Context, id string, p *model.Patient) error {
	r.m.Lock()
	defer r.m.Unlock()
	patient, ok := r.patients[id]
	if !ok {
		return fmt.Errorf("patient with id %s does not exist", id)
	}
	res := converter.ToUpdate(&patient, *p)

	r.patients[id] = *res

	err := r.saveInFile()
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DelPatient(_ context.Context, id string) error {
	r.m.Lock()
	defer r.m.Unlock()
	_, ok := r.patients[id]
	if !ok {
		return fmt.Errorf("patient with id %s does not exist", id)
	}
	delete(r.patients, id)

	err := r.saveInFile()

	if err != nil {
		return err
	}

	return nil
}

func (r *repo) saveInFile() (err error) {

	ps := make([]datamodel.Patient, 0, len(r.patients))
	for _, v := range r.patients {
		ps = append(ps, v)
	}
	data, err := json.Marshal(ps)

	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, fs.ModeAppend)
	if err != nil {
		return err
	}

	return nil

}
