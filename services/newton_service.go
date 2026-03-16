// services/newton_service.go
package services

import (
	"errors"
	"math"

	"lab-cuantitativo/models"
)

// CalculateNewton resuelve f(x)=0 usando el método de Newton-Raphson
func CalculateNewton(req models.NewtonRequest) (models.NewtonResponse, error) {
	tol := req.Tol
	if tol == 0 {
		tol = 1e-6
	}
	maxIter := req.MaxIter
	if maxIter == 0 {
		maxIter = 50
	}

	// Validar que la expresión sea evaluable
	if _, err := EvalExpr(req.Expression, req.X0); err != nil {
		return models.NewtonResponse{}, errors.New("expresión inválida: " + err.Error())
	}

	x := req.X0
	var tabla []models.NewtonRow
	converged := false

	for i := 0; i < maxIter; i++ {
		fx, _ := EvalExpr(req.Expression, x)
		dfx, _ := Derivative(req.Expression, x)

		if math.Abs(dfx) < 1e-14 {
			return models.NewtonResponse{}, errors.New("derivada cero — el método diverge en este punto")
		}

		xNew := x - fx/dfx
		errVal := math.Abs(xNew - x)

		tabla = append(tabla, models.NewtonRow{
			Iteracion: i + 1,
			X:         Round(x, 8),
			Fx:        Round(fx, 8),
			Dfx:       Round(dfx, 8),
			XNew:      Round(xNew, 8),
			Error:     errVal,
		})

		x = xNew
		if errVal < tol {
			converged = true
			break
		}
	}

	fRaiz, _ := EvalExpr(req.Expression, x)

	return models.NewtonResponse{
		Raiz:        Round(x, 10),
		FRaiz:       fRaiz,
		Iteraciones: len(tabla),
		Convergido:  converged,
		Tabla:       tabla,
	}, nil
}