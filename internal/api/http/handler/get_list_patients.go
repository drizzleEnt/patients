package handler

import (
	"net/http"

	"github.com/drizzleent/patients/internal/api"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetListPatients(ctx *gin.Context) {
	res, err := h.service.GetListPatients(ctx)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, res)
}
