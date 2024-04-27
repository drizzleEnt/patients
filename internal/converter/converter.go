package converter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/drizzleent/patients/internal/model"
	"github.com/gin-gonic/gin"
)

const (
	gender = "gender"
	create = "create"
	id     = "id"
)

func FromReqToPatient(c *gin.Context) (*model.Patient, int, error) {
	var patient model.Patient
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	err = json.Unmarshal(body, &patient)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if strings.Contains(c.Request.URL.Path, create) {
		if patient.Fullname == "" || patient.Birthday == "" {
			return nil, http.StatusBadRequest, errors.New("patient name or birthday requared")
		}
	}

	if strings.Contains(string(body), gender) {
		patient.IsGenderValid = true
	}
	return &patient, http.StatusOK, nil
}

func FromReqToId(c *gin.Context) (string, int, error) {
	key := c.Param(id)
	if len(key) == 0 {
		return "", http.StatusBadRequest, errors.New("patient id is requared")
	}

	return key, http.StatusOK, nil
}

func FromPatientToReq(p *model.Patient) *model.ReqPatient {
	return &model.ReqPatient{
		Fullname: p.Fullname,
		Birthday: p.Birthday,
		Gender:   p.Gender,
		Guid:     p.Guid,
	}
}
