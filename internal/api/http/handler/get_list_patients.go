package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetListPatients(ctx *gin.Context) {
	fmt.Println("get")
}
