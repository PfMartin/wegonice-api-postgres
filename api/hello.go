package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello")
}
