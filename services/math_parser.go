// services/math_parser.go
package services

import (
	"errors"
	"math"
)

// EvalExpr evalúa una expresión matemática con una variable x
// Soporta: +, -, *, /, ** (potencia), sin, cos, tan, exp, log, sqrt, abs, pi, e
func EvalExpr(expr string, x float64) (float64, error) {
	p := &parser{input: expr, pos: 0, vars: map[string]float64{"x": x}}
	result, err := p.parseExpr()
	if err != nil {
		return 0, err
	}
	return result, nil
}

// Derivative calcula la derivada numérica usando diferencias centradas
func Derivative(expr string, x float64) (float64, error) {
	h := 1e-7
	f1, err := EvalExpr(expr, x+h)
	if err != nil {
		return 0, err
	}
	f2, err := EvalExpr(expr, x-h)
	if err != nil {
		return 0, err
	}
	return (f1 - f2) / (2 * h), nil
}

// Round redondea a n decimales
func Round(val float64, decimals int) float64 {
	pow := math.Pow(10, float64(decimals))
	return math.Round(val*pow) / pow
}

// ── Parser interno ───────────────────────────────────────────

type parser struct {
	input string
	pos   int
	vars  map[string]float64
}

func (p *parser) skipSpaces() {
	for p.pos < len(p.input) && (p.input[p.pos] == ' ' || p.input[p.pos] == '\t') {
		p.pos++
	}
}

func (p *parser) peek() byte {
	p.skipSpaces()
	if p.pos >= len(p.input) {
		return 0
	}
	return p.input[p.pos]
}

func (p *parser) parseExpr() (float64, error) {
	return p.parseAddSub()
}

func (p *parser) parseAddSub() (float64, error) {
	left, err := p.parseMulDiv()
	if err != nil {
		return 0, err
	}
	for {
		c := p.peek()
		if c == '+' {
			p.pos++
			right, err := p.parseMulDiv()
			if err != nil {
				return 0, err
			}
			left += right
		} else if c == '-' {
			p.pos++
			right, err := p.parseMulDiv()
			if err != nil {
				return 0, err
			}
			left -= right
		} else {
			break
		}
	}
	return left, nil
}

func (p *parser) parseMulDiv() (float64, error) {
	left, err := p.parsePow()
	if err != nil {
		return 0, err
	}
	for {
		p.skipSpaces()
		if p.pos < len(p.input) && p.input[p.pos] == '*' &&
			p.pos+1 < len(p.input) && p.input[p.pos+1] != '*' {
			p.pos++
			right, err := p.parsePow()
			if err != nil {
				return 0, err
			}
			left *= right
		} else if p.peek() == '/' {
			p.pos++
			right, err := p.parsePow()
			if err != nil {
				return 0, err
			}
			if right == 0 {
				return 0, errors.New("división por cero")
			}
			left /= right
		} else {
			break
		}
	}
	return left, nil
}

func (p *parser) parsePow() (float64, error) {
	base, err := p.parseUnary()
	if err != nil {
		return 0, err
	}
	p.skipSpaces()
	if p.pos+1 < len(p.input) && p.input[p.pos] == '*' && p.input[p.pos+1] == '*' {
		p.pos += 2
		exp, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		return math.Pow(base, exp), nil
	}
	return base, nil
}

func (p *parser) parseUnary() (float64, error) {
	p.skipSpaces()
	if p.peek() == '-' {
		p.pos++
		val, err := p.parsePrimary()
		if err != nil {
			return 0, err
		}
		return -val, nil
	}
	return p.parsePrimary()
}

func (p *parser) parsePrimary() (float64, error) {
	p.skipSpaces()
	c := p.peek()

	if c >= '0' && c <= '9' || c == '.' {
		return p.parseNumber()
	}
	if c == '(' {
		p.pos++
		val, err := p.parseExpr()
		if err != nil {
			return 0, err
		}
		p.skipSpaces()
		if p.pos < len(p.input) && p.input[p.pos] == ')' {
			p.pos++
		}
		return val, nil
	}
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' {
		return p.parseIdentifier()
	}
	return 0, nil
}

func (p *parser) parseNumber() (float64, error) {
	start := p.pos
	for p.pos < len(p.input) {
		ch := p.input[p.pos]
		if ch >= '0' && ch <= '9' || ch == '.' {
			p.pos++
		} else if (ch == 'e' || ch == 'E') && p.pos > start {
			p.pos++
			if p.pos < len(p.input) && (p.input[p.pos] == '+' || p.input[p.pos] == '-') {
				p.pos++
			}
		} else {
			break
		}
	}
	val := 0.0
	s := p.input[start:p.pos]
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			val = val*10 + float64(ch-'0')
		}
	}
	// Parseado simple — usar strconv para mayor precisión
	import_val, err := parseFloatSimple(p.input[start:p.pos])
	if err != nil {
		return val, nil
	}
	return import_val, nil
}

func parseFloatSimple(s string) (float64, error) {
	if len(s) == 0 {
		return 0, errors.New("número vacío")
	}
	val := 0.0
	frac := 0.0
	fracDiv := 1.0
	inFrac := false
	expPart := 0
	hasExp := false
	expSign := 1

	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			if hasExp {
				expPart = expPart*10 + int(ch-'0')
			} else if inFrac {
				frac = frac*10 + float64(ch-'0')
				fracDiv *= 10
			} else {
				val = val*10 + float64(ch-'0')
			}
		} else if ch == '.' && !inFrac && !hasExp {
			inFrac = true
		} else if (ch == 'e' || ch == 'E') && !hasExp {
			hasExp = true
			_ = i
		} else if ch == '+' && hasExp && expPart == 0 {
			expSign = 1
		} else if ch == '-' && hasExp && expPart == 0 {
			expSign = -1
		}
	}

	result := val + frac/fracDiv
	if hasExp {
		result *= math.Pow(10, float64(expSign*expPart))
	}
	return result, nil
}

func (p *parser) parseIdentifier() (float64, error) {
	start := p.pos
	for p.pos < len(p.input) {
		ch := p.input[p.pos]
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' ||
			ch >= '0' && ch <= '9' || ch == '_' {
			p.pos++
		} else {
			break
		}
	}
	name := p.input[start:p.pos]
	p.skipSpaces()

	if p.pos < len(p.input) && p.input[p.pos] == '(' {
		p.pos++
		arg, err := p.parseExpr()
		if err != nil {
			return 0, err
		}
		p.skipSpaces()
		if p.pos < len(p.input) && p.input[p.pos] == ')' {
			p.pos++
		}
		switch name {
		case "sin":
			return math.Sin(arg), nil
		case "cos":
			return math.Cos(arg), nil
		case "tan":
			return math.Tan(arg), nil
		case "exp":
			return math.Exp(arg), nil
		case "log", "ln":
			if arg <= 0 {
				return 0, errors.New("log de número no positivo")
			}
			return math.Log(arg), nil
		case "log10":
			return math.Log10(arg), nil
		case "sqrt":
			if arg < 0 {
				return 0, errors.New("sqrt de número negativo")
			}
			return math.Sqrt(arg), nil
		case "abs":
			return math.Abs(arg), nil
		default:
			return 0, errors.New("función desconocida: " + name)
		}
	}

	// Variables y constantes
	if val, ok := p.vars[name]; ok {
		return val, nil
	}
	switch name {
	case "pi", "PI":
		return math.Pi, nil
	case "e":
		return math.E, nil
	}
	return 0, errors.New("variable desconocida: " + name)
}