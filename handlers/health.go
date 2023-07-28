package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func (hc *HealthHandler) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
