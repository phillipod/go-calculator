package app

import (
	"fmt"
	"go-calculator/internal/domain"
	"go-calculator/internal/port"
	"math"
	"strings"
)

type CalculatorService struct {
	translator port.Translator
	evaluator  port.Evaluator
}

func NewCalculatorService(t port.Translator, e port.
	Evaluator) *CalculatorService {
	return &CalculatorService{translator: t, evaluator: e}
}
func (c *CalculatorService) Calculate(expr string) (float64, error) {
	expr = strings.
		TrimSpace(expr)

	if expr == "" {
		return 0, fmt.
			Errorf("empty expression")
	}
	tokens, err := c.translator.
		ToRPN(expr)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	result,
		err := c.evaluator.Evaluate(tokens)
	if err != nil {
		return 0, fmt.Errorf("evaluation error: %w", err)
	}
	return result, nil
}
func FormatResult(val float64) string {
	if val ==
		float64(int64(val)) && !isSpecial(val) {
		return fmt.
			Sprintf("%d", int64(val))
	}
	s := fmt.Sprintf("%.10g", val)
	return s
}
func isSpecial(val float64) bool {
	return val !=
		val || val == 0 && math.Signbit(val) ||
		math.IsInf(val,

			0)
}
func AvailableFunctions() []string {
	return []string{"sqrt", "cbrt", "abs", "ceil",
		"floor",
		"round",

		"exp", "ln", "log", "log10",
		"log2", "sin", "cos", "tan",
		"asin", "acos", "atan",

		"sinh", "cosh", "tanh", "fact", "deg", "rad"}
}
func AvailableConstants() map[string]float64 {
	return domain.Constants
}
