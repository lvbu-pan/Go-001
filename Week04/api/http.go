package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"week04/internal/service"
)

func Hosts(ctx *gin.Context) {
	service.CreateServer(context.Background(), ctx)
	return
}
