// handlers/lab_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lab-cuantitativo/models"
	"lab-cuantitativo/services"
)

// VPNTIRHandler maneja POST /api/lab/vpn-tir/
func VPNTIRHandler(c *gin.Context) {
	var req models.VPNRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := services.CalculateVPN(req)
	c.JSON(http.StatusOK, result)
}

// AnualidadesHandler maneja POST /api/lab/anualidades/
func AnualidadesHandler(c *gin.Context) {
	var req models.AnualidadesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := services.CalculateAnualidades(req)
	c.JSON(http.StatusOK, result)
}

// SimplexHandler maneja POST /api/lab/simplex/
func SimplexHandler(c *gin.Context) {
	var req models.SimplexRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CalculateSimplex(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// EstadisticaHandler maneja POST /api/lab/estadistica/
func EstadisticaHandler(c *gin.Context) {
	var req models.EstadisticaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := services.CalculateEstadistica(req)
	c.JSON(http.StatusOK, result)
}

// NewtonHandler maneja POST /api/lab/newton/
func NewtonHandler(c *gin.Context) {
	var req models.NewtonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CalculateNewton(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// IntegracionHandler maneja POST /api/lab/integracion/
func IntegracionHandler(c *gin.Context) {
	var req models.IntegracionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CalculateIntegracion(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}