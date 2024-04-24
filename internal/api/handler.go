package api

import "github.com/gin-gonic/gin"

type Handler interface {
	GetListPatients(*gin.Context)
	NewPatient(*gin.Context)
	EditPatient(*gin.Context)
	DelPatient(*gin.Context)
	InitRoutes() *gin.Engine
}
