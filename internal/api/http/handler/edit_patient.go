package handler

import (
	"net/http"

	"github.com/drizzleent/patients/internal/api"
	"github.com/drizzleent/patients/internal/converter"
	"github.com/gin-gonic/gin"
)

func (h *handler) EditPatient(ctx *gin.Context) {
	id, status, err := converter.FromReqToId(ctx)
	if err != nil {
		api.NewErrorResponse(ctx, status, err.Error())
		return
	}

	p, status, err := converter.FromReqToPatient(ctx)

	if err != nil {
		api.NewErrorResponse(ctx, status, err.Error())
		return
	}

	err = h.service.EditPatient(ctx, id, p)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "updated")
}
