package handler

import "github.com/gin-gonic/gin"

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
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
