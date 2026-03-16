// services/simplex_service.go
package services

import (
	"errors"
	"math"

	"lab-cuantitativo/models"
)

// CalculateSimplex resuelve un problema de PL por el método Simplex
func CalculateSimplex(req models.SimplexRequest) (models.SimplexResponse, error) {
	c := req.C
	A := req.A
	b := req.B
	m := len(A)
	n := len(c)
	cols := n + m + 1

	// Construir tableau con variables de holgura
	tableau := make([][]float64, m+1)
	for i := 0; i <= m; i++ {
		tableau[i] = make([]float64, cols)
	}
	for i := 0; i < m; i++ {
		copy(tableau[i], A[i])
		tableau[i][n+i] = 1
		tableau[i][cols-1] = b[i]
	}
	for j := 0; j < n; j++ {
		tableau[m][j] = -c[j]
	}

	var iterations []models.SimplexIter

	for iter := 0; iter < 100; iter++ {
		// Columna pivote: coeficiente más negativo en fila objetivo
		pivCol := -1
		minVal := -1e-9
		for j := 0; j < cols-1; j++ {
			if tableau[m][j] < minVal {
				minVal = tableau[m][j]
				pivCol = j
			}
		}
		if pivCol == -1 {
			break // Óptimo encontrado
		}

		// Fila pivote: razón mínima positiva
		pivRow := -1
		minRatio := math.MaxFloat64
		for i := 0; i < m; i++ {
			if tableau[i][pivCol] > 1e-9 {
				ratio := tableau[i][cols-1] / tableau[i][pivCol]
				if ratio < minRatio {
					minRatio = ratio
					pivRow = i
				}
			}
		}
		if pivRow == -1 {
			return models.SimplexResponse{}, errors.New("problema no acotado")
		}

		iterations = append(iterations, models.SimplexIter{
			Iteration: iter + 1,
			PivotCol:  pivCol,
			PivotRow:  pivRow,
		})

		// Operación de pivote (eliminación gaussiana)
		piv := tableau[pivRow][pivCol]
		for j := 0; j < cols; j++ {
			tableau[pivRow][j] /= piv
		}
		for i := 0; i <= m; i++ {
			if i != pivRow {
				factor := tableau[i][pivCol]
				for j := 0; j < cols; j++ {
					tableau[i][j] -= factor * tableau[pivRow][j]
				}
			}
		}
	}

	// Extraer solución óptima
	solution := make(map[string]float64)
	for j := 0; j < n; j++ {
		ones, zeros := 0, 0
		rowIdx := -1
		for i := 0; i < m; i++ {
			if math.Abs(tableau[i][j]-1) < 1e-9 {
				ones++
				rowIdx = i
			} else if math.Abs(tableau[i][j]) < 1e-9 {
				zeros++
			}
		}
		varName := "x" + string(rune('0'+j+1))
		if ones == 1 && zeros == m-1 && rowIdx >= 0 {
			solution[varName] = Round(tableau[rowIdx][cols-1], 6)
		} else {
			solution[varName] = 0
		}
	}

	z := Round(-tableau[m][cols-1], 6)

	return models.SimplexResponse{
		Solucion:    solution,
		ZOptimo:     z,
		Iteraciones: iterations,
		Factible:    true,
	}, nil
}