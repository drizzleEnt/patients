package handler

import (
	"github.com/drizzleent/patients/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.ApiService
}

func NewHandler(srv service.ApiService) *handler {
	return &handler{
		service: srv,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/")
	{
		api.GET("/", h.GetListPatients)
		api.POST("/create", h.NewPatient)
		api.POST("/edit", h.EditPatient)
		api.POST("/delete", h.DelPatient)

	}

	return router
}
