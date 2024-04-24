package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *handler) NewPatient(ctx *gin.Context) {
	fmt.Println("new")
}
