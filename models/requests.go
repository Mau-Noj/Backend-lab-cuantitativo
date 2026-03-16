// models/requests.go
package models

// VPNRequest representa la solicitud para calcular VPN y TIR
type VPNRequest struct {
	Rate  float64   `json:"rate"  binding:"required"`
	Flows []float64 `json:"flows" binding:"required,min=1"`
}

// AnualidadesRequest representa la solicitud para calcular anualidades
type AnualidadesRequest struct {
	A float64 `json:"A" binding:"required"`
	I float64 `json:"i" binding:"required"`
	N int     `json:"n" binding:"required,min=1"`
}

// SimplexRequest representa la solicitud para el método Simplex
type SimplexRequest struct {
	C []float64   `json:"c" binding:"required,min=1"`
	A [][]float64 `json:"A" binding:"required,min=1"`
	B []float64   `json:"b" binding:"required,min=1"`
}

// EstadisticaRequest representa la solicitud para estadística descriptiva
type EstadisticaRequest struct {
	Data []float64 `json:"data" binding:"required,min=2"`
}

// NewtonRequest representa la solicitud para Newton-Raphson
type NewtonRequest struct {
	Expression string  `json:"expression" binding:"required"`
	X0         float64 `json:"x0"`
	Tol        float64 `json:"tol"`
	MaxIter    int     `json:"max_iter"`
}

// IntegracionRequest representa la solicitud para integración numérica
type IntegracionRequest struct {
	Expression string  `json:"expression" binding:"required"`
	A          float64 `json:"a"`
	B          float64 `json:"b"`
	N          int     `json:"n"`
}