package converter

import (
	"errors"
	"net/http"

	"github.com/drizzleent/patients/internal/model"
	"github.com/gin-gonic/gin"
)

func FromReqToPatient(c *gin.Context) (*model.Patient, int, error) {
	var patient model.Patient

	err := c.BindJSON(&patient)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return &patient, http.StatusOK, nil
}

func FromReqToId(c *gin.Context) (string, int, error) {
	key := c.Param("id")
	if len(key) == 0 {
		return "", http.StatusBadRequest, errors.New("patient id is requared")
	}

	return key, http.StatusOK, nil
}
