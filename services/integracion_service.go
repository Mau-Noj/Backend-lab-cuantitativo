// services/integracion_service.go
package services

import (
	"errors"
	"math"

	"lab-cuantitativo/models"
)

// CalculateIntegracion aproxima ∫f(x)dx por Trapecio y Simpson 1/3
func CalculateIntegracion(req models.IntegracionRequest) (models.IntegracionResponse, error) {
	n := req.N
	if n == 0 {
		n = 100
	}
	if n%2 != 0 {
		n++ // Simpson requiere n par
	}

	a, b := req.A, req.B
	h := (b - a) / float64(n)

	// Validar expresión
	if _, err := EvalExpr(req.Expression, a); err != nil {
		return models.IntegracionResponse{}, errors.New("expresión inválida: " + err.Error())
	}

	f := func(x float64) float64 {
		v, _ := EvalExpr(req.Expression, x)
		return v
	}

	// Método del Trapecio
	trap := (f(a) + f(b)) / 2
	for i := 1; i < n; i++ {
		trap += f(a + float64(i)*h)
	}
	trap *= h

	// Método de Simpson 1/3
	simp := f(a) + f(b)
	for i := 1; i < n; i++ {
		xi := a + float64(i)*h
		if i%2 != 0 {
			simp += 4 * f(xi)
		} else {
			simp += 2 * f(xi)
		}
	}
	simp *= h / 3

	// Puntos para graficar (máx 20 puntos)
	step := n / 20
	if step < 1 {
		step = 1
	}
	var puntos []models.IntegracionPoint
	for i := 0; i <= n; i += step {
		xi := a + float64(i)*h
		puntos = append(puntos, models.IntegracionPoint{
			X: Round(xi, 4),
			Y: Round(f(xi), 6),
		})
	}

	return models.IntegracionResponse{
		Trapecio:   Round(trap, 10),
		Simpson:    Round(simp, 10),
		Diferencia: Round(math.Abs(simp-trap), 10),
		NUsado:     n,
		H:          Round(h, 6),
		Puntos:     puntos,
	}, nil
}