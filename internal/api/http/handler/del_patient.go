package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *handler) DelPatient(ctx *gin.Context) {
	fmt.Println("del")
}
