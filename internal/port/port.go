package port

import "go-calculator/internal/domain"

type Translator interface {
	ToRPN(expr string) ([]domain.Token, error)
}
type Evaluator interface {
	Evaluate(tokens []domain.Token) (float64,
		error)
}
type Calculator interface {
	Calculate(expr string) (float64, error)
}
