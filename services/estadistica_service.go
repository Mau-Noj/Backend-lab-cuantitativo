// services/estadistica_service.go
package services

import (
	"math"
	"sort"

	"lab-cuantitativo/models"
)

// CalculateEstadistica calcula las medidas descriptivas de un conjunto de datos
func CalculateEstadistica(req models.EstadisticaRequest) models.EstadisticaResponse {
	data := make([]float64, len(req.Data))
	copy(data, req.Data)
	sort.Float64s(data)
	n := len(data)

	// Media
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	mean := sum / float64(n)

	// Mediana
	var median float64
	if n%2 == 0 {
		median = (data[n/2-1] + data[n/2]) / 2
	} else {
		median = data[n/2]
	}

	// Moda (solo si algún valor se repite)
	freq := make(map[float64]int)
	for _, v := range data {
		freq[v]++
	}
	maxFreq := 0
	var modeVal float64
	for v, f := range freq {
		if f > maxFreq {
			maxFreq = f
			modeVal = v
		}
	}
	var mode *float64
	if maxFreq > 1 {
		mode = &modeVal
	}

	// Varianza muestral y desviación estándar
	variance := 0.0
	for _, v := range data {
		variance += (v - mean) * (v - mean)
	}
	variance /= float64(n - 1)
	std := math.Sqrt(variance)

	// Coeficiente de variación
	var cv *float64
	if mean != 0 {
		cvVal := Round((std/mean)*100, 4)
		cv = &cvVal
	}

	// Cuartiles
	q1 := data[int(math.Floor(float64(n)*0.25))]
	q3 := data[int(math.Floor(float64(n)*0.75))]
	iqr := q3 - q1

	return models.EstadisticaResponse{
		N:                 n,
		Media:             Round(mean, 6),
		Mediana:           Round(median, 6),
		Moda:              mode,
		Varianza:          Round(variance, 6),
		DesvEst:           Round(std, 6),
		CV:                cv,
		Minimo:            data[0],
		Maximo:            data[n-1],
		Q1:                q1,
		Q3:                q3,
		RangoIntercuartil: Round(iqr, 6),
		DatosOrdenados:    data,
	}
}