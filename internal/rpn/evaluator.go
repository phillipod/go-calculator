package rpn

import (
	"fmt"
	"go-calculator/internal/domain"
	"math"
)

type StackEvaluator struct{}

func NewStackEvaluator() *StackEvaluator {
	return &StackEvaluator{}
}
func (e *StackEvaluator) Evaluate(tokens []domain.
	Token) (
	float64, error) {
	stack := make(
		[]float64, 0, len(tokens))
	for _, tok := range tokens {
		switch tok.
			Type {
		case domain.TokenNumber, domain.TokenConstant:
			stack = append(stack, tok.NumValue)
		case domain.TokenOperator:
			info, ok := domain.GetOpInfo(tok.Value)
			if !ok {
				return 0, fmt.Errorf("unknown operator: %s", tok.Value)
			}
			switch info.Arity {
			case domain.Binary:
				if len(stack) < 2 {
					return 0, fmt.Errorf("not enough operands for %s", tok.Value)
				}
				b := stack[len(stack)-1]
				a := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				result, err := applyBinary(tok.Value, a, b)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
			case domain.Unary:
				if len(stack) < 1 {
					return 0, fmt.Errorf("not enough operands for %s", tok.Value)
				}
				a := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result, err := applyUnary(tok.Value, a)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
			}
		case domain.TokenFunction:
			info, ok := domain.GetFuncInfo(tok.Value)
			if !ok {
				return 0, fmt.Errorf("unknown function: %s", tok.Value)
			}
			if len(stack) < int(info.Arity) {
				return 0, fmt.Errorf("not enough arguments for %s", tok.Value)
			}
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, info.Apply(a))
		default:
			return 0, fmt.Errorf("unexpected token type in RPN: %v", tok.Type)
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("expression incomplete: stack has %d items", len(stack))
	}
	return stack[0], nil
}
func applyBinary(
	op string, a, b float64) (float64, error) {

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a *
				b,
			nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, fmt.Errorf("modulo by zero")
		}
		return math.
			Mod(a, b), nil
	case "^":
		return math.Pow(a,
			b), nil
	default:
		return 0, fmt.Errorf("unknown binary operator: %s",
			op)
	}
}
func applyUnary(op string, a float64) (float64, error) {
	switch op {
	case "neg":
		return -a, nil
	default:
		return 0, fmt.Errorf("unknown unary operator: %s",

			op)
	}
}
