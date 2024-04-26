package handler

import (
	"net/http"

	"github.com/drizzleent/patients/internal/api"
	"github.com/drizzleent/patients/internal/converter"
	"github.com/gin-gonic/gin"
)

func (h *handler) NewPatient(ctx *gin.Context) {
	req, status, err := converter.FromReqToPatient(ctx)
	if err != nil {
		api.NewErrorResponse(ctx, status, err.Error())
		return
	}

	id, err := h.service.NewPatient(ctx, req)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
