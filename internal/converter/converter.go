package converter

import (
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
