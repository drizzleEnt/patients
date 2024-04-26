package handler

import (
	"net/http"

	"github.com/drizzleent/patients/internal/api"
	"github.com/drizzleent/patients/internal/converter"
	"github.com/gin-gonic/gin"
)

func (h *handler) DelPatient(ctx *gin.Context) {
	id, status, err := converter.FromReqToId(ctx)
	if err != nil {
		api.NewErrorResponse(ctx, status, err.Error())
		return
	}

	err = h.service.DelPatient(ctx, id)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
