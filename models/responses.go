// models/responses.go
package models

// ── VPN & TIR ────────────────────────────────────────────────

type VPNStep struct {
	Step  int     `json:"step"`
	Flujo float64 `json:"flujo"`
	VP    float64 `json:"vp"`
}

type VPNResponse struct {
	VPN           float64   `json:"vpn"`
	TIR           float64   `json:"tir"`
	Viable        bool      `json:"viable"`
	Decision      string    `json:"decision"`
	Steps         []VPNStep `json:"steps"`
	IteracionesTIR int      `json:"iteraciones_tir"`
}

// ── Anualidades ──────────────────────────────────────────────

type AmortizacionRow struct {
	Periodo      int     `json:"periodo"`
	SaldoInicial float64 `json:"saldo_inicial"`
	Interes      float64 `json:"interes"`
	Pago         float64 `json:"pago"`
}

type AnualidadesResponse struct {
	VP                float64           `json:"vp"`
	VF                float64           `json:"vf"`
	TotalPagado       float64           `json:"total_pagado"`
	TotalIntereses    float64           `json:"total_intereses"`
	TablaAmortizacion []AmortizacionRow `json:"tabla_amortizacion"`
}

// ── Simplex ──────────────────────────────────────────────────

type SimplexIter struct {
	Iteration int `json:"iteration"`
	PivotCol  int `json:"pivot_col"`
	PivotRow  int `json:"pivot_row"`
}

type SimplexResponse struct {
	Solucion    map[string]float64 `json:"solucion"`
	ZOptimo     float64            `json:"z_optimo"`
	Iteraciones []SimplexIter      `json:"iteraciones"`
	Factible    bool               `json:"factible"`
}

// ── Estadística ──────────────────────────────────────────────

type EstadisticaResponse struct {
	N                 int      `json:"n"`
	Media             float64  `json:"media"`
	Mediana           float64  `json:"mediana"`
	Moda              *float64 `json:"moda"`
	Varianza          float64  `json:"varianza"`
	DesvEst           float64  `json:"desv_est"`
	CV                *float64 `json:"cv"`
	Minimo            float64  `json:"minimo"`
	Maximo            float64  `json:"maximo"`
	Q1                float64  `json:"q1"`
	Q3                float64  `json:"q3"`
	RangoIntercuartil float64  `json:"rango_intercuartil"`
	DatosOrdenados    []float64 `json:"datos_ordenados"`
}

// ── Newton-Raphson ───────────────────────────────────────────

type NewtonRow struct {
	Iteracion int     `json:"iteracion"`
	X         float64 `json:"x"`
	Fx        float64 `json:"fx"`
	Dfx       float64 `json:"dfx"`
	XNew      float64 `json:"x_new"`
	Error     float64 `json:"error"`
}

type NewtonResponse struct {
	Raiz        float64     `json:"raiz"`
	FRaiz       float64     `json:"f_raiz"`
	Iteraciones int         `json:"iteraciones"`
	Convergido  bool        `json:"convergido"`
	Tabla       []NewtonRow `json:"tabla"`
}

// ── Integración ──────────────────────────────────────────────

type IntegracionPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type IntegracionResponse struct {
	Trapecio   float64            `json:"trapecio"`
	Simpson    float64            `json:"simpson"`
	Diferencia float64            `json:"diferencia"`
	NUsado     int                `json:"n_usado"`
	H          float64            `json:"h"`
	Puntos     []IntegracionPoint `json:"puntos"`
}