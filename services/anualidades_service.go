// services/anualidades_service.go
package services

import (
	"math"

	"lab-cuantitativo/models"
)

// CalculateAnualidades calcula VP, VF y tabla de amortización
func CalculateAnualidades(req models.AnualidadesRequest) models.AnualidadesResponse {
	A := req.A
	i := req.I / 100
	n := req.N

	vp := A * (1 - math.Pow(1+i, float64(-n))) / i
	vf := A * (math.Pow(1+i, float64(n)) - 1) / i
	totalPagado := A * float64(n)
	totalIntereses := vf - totalPagado

	// Tabla de amortización (máx 12 períodos visibles)
	limit := n
	if limit > 12 {
		limit = 12
	}

	tabla := make([]models.AmortizacionRow, limit)
	saldo := vp
	for t := 1; t <= limit; t++ {
		interes := saldo * i
		capital := A - interes
		tabla[t-1] = models.AmortizacionRow{
			Periodo:      t,
			SaldoInicial: Round(saldo, 2),
			Interes:      Round(interes, 2),
			Pago:         Round(A, 2),
		}
		saldo -= capital
	}

	return models.AnualidadesResponse{
		VP:                Round(vp, 4),
		VF:                Round(vf, 4),
		TotalPagado:       Round(totalPagado, 2),
		TotalIntereses:    Round(totalIntereses, 2),
		TablaAmortizacion: tabla,
	}
}