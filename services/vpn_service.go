// services/vpn_service.go
package services

import (
	"math"

	"lab-cuantitativo/models"
)

// CalculateVPN calcula el Valor Presente Neto y la Tasa Interna de Retorno
func CalculateVPN(req models.VPNRequest) models.VPNResponse {
	rate := req.Rate / 100
	flows := req.Flows

	// VPN
	vpn := 0.0
	for t, f := range flows {
		vpn += f / math.Pow(1+rate, float64(t))
	}

	// TIR — Newton-Raphson sobre la función VPN(r) = 0
	tir := rate
	if tir <= 0 {
		tir = 0.1
	}
	iters := 0
	for i := 0; i < 1000; i++ {
		fVal := 0.0
		dfVal := 0.0
		for t, f := range flows {
			fVal += f / math.Pow(1+tir, float64(t))
			if t > 0 {
				dfVal -= float64(t) * f / math.Pow(1+tir, float64(t+1))
			}
		}
		if math.Abs(dfVal) < 1e-12 {
			break
		}
		newTir := tir - fVal/dfVal
		if math.Abs(newTir-tir) < 1e-8 {
			tir = newTir
			break
		}
		tir = newTir
		iters++
	}

	// Tabla de pasos
	steps := make([]models.VPNStep, len(flows))
	for t, f := range flows {
		vp := f / math.Pow(1+rate, float64(t))
		steps[t] = models.VPNStep{
			Step:  t,
			Flujo: Round(f, 2),
			VP:    Round(vp, 4),
		}
	}

	decision := "Rechazar"
	if vpn >= 0 {
		decision = "Aceptar"
	}

	return models.VPNResponse{
		VPN:            Round(vpn, 4),
		TIR:            Round(tir*100, 4),
		Viable:         vpn >= 0,
		Decision:       decision,
		Steps:          steps,
		IteracionesTIR: iters,
	}
}