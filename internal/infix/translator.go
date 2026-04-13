package infix

import (
	"fmt"
	"go-calculator/internal/domain"
	"strconv"
	"strings"
	"unicode"
)

type ShuntingYard struct{}

func NewShuntingYard() *ShuntingYard {
	return &ShuntingYard{}
}
func (s *ShuntingYard) ToRPN(expr string) ([]domain.
	Token, error) {
	tokens,
		err :=
		s.
			tokenize(expr)
	if err != nil {
		return nil, err
	}
	return s.shunt(tokens)
}
func (s *ShuntingYard) tokenize(expr string) ([]domain.
	Token,
	error) {
	var tokens []domain.
		Token
	runes := []rune(expr)
	i := 0
	for i < len(runes) {
		ch := runes[i]
		if unicode.
			IsSpace(ch) {
			i++
			continue
		}
		if ch ==
			'(' {
			tokens =
				append(tokens,
					domain.
						NewLeftParenToken())
			i++
			continue
		}
		if ch ==
			')' {
			tokens = append(tokens, domain.NewRightParenToken())
			i++
			continue
		}
		if isDigitOrDot(ch) {
			j := i
			dotSeen := false
			for j < len(runes) && (isDigitOrDot(runes[j]) || runes[j] == 'e' || runes[j] == 'E' || (j > 0 && (runes[j-1] == 'e' || runes[j-1] == 'E') && (runes[j] == '+' || runes[j] == '-'))) {
				if runes[j] == '.' {
					if dotSeen {
						return nil, fmt.Errorf("invalid number at position %d", i)
					}
					dotSeen = true
				}
				j++
			}
			numStr := string(runes[i:j])
			val, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", numStr)
			}
			tokens = append(tokens, domain.NewNumberToken(val))
			i = j
			continue
		}
		if isAlphaUnderscore(ch) {
			j := i
			for j < len(runes) && isAlphaUnderscore(runes[j]) {
				j++
			}
			name := strings.ToLower(string(runes[i:j]))
			if val, ok := domain.GetConstant(name); ok {
				tokens = append(tokens, domain.NewConstantToken(name, val))
			} else if domain.IsFunction(name) {
				tokens = append(tokens, domain.NewFunctionToken(name))
			} else {
				return nil, fmt.Errorf("unknown identifier: %s", name)
			}
			i = j
			continue
		}
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '%' || ch == '^' {
			op := string(ch)
			if ch == '-' && isUnaryContext(tokens) {
				op = "neg"
			}
			tokens = append(tokens, domain.NewOperatorToken(op))
			i++
			continue
		}
		return nil, fmt.Errorf("unexpected character at position %d: %c", i, ch)
	}
	return tokens, nil
}
func isUnaryContext(tokens []domain.Token) bool {
	if len(tokens) == 0 {
		return true

	}
	last := tokens[len(tokens)-1]
	return last.Type == domain.TokenOperator ||
		last.Type ==

			domain.TokenLeftParen
}
func (s *ShuntingYard) shunt(tokens []domain.Token) ([]domain.
	Token, error) {
	var output []domain.Token
	var opStack []domain.Token
	for _, tok := range tokens {
		switch tok.
			Type {
		case domain.TokenNumber, domain.TokenConstant:
			output = append(output,
				tok)
		case
			domain.TokenFunction:
			opStack =
				append(opStack, tok)
		case domain.TokenOperator:
			for len(opStack) > 0 {
				top := opStack[len(opStack)-1]
				if top.
					Type == domain.TokenLeftParen {
					break
				}
				if shouldPop(top, tok) {
					output = append(output, top)
					opStack = opStack[:len(opStack)-1]
				} else {
					break
				}
			}
			opStack = append(opStack,
				tok)
		case domain.TokenLeftParen:
			opStack =
				append(opStack,
					tok)
		case domain.TokenRightParen:
			foundLP := false
			for len(opStack) >
				0 {
				top := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				if top.Type == domain.TokenLeftParen {
					foundLP = true
					break
				}
				output = append(output, top)
			}
			if !foundLP {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			if len(opStack) > 0 && opStack[len(opStack)-1].Type == domain.TokenFunction {
				output = append(output,
					opStack[len(
						opStack,
					)-1])
				opStack = opStack[:len(opStack)-1]
			}
		default:
			return nil,
				fmt.Errorf("unexpected token: %v",
					tok)
		}
	}
	for len(opStack) > 0 {
		top :=
			opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		if top.Type ==
			domain.
				TokenLeftParen || top.Type ==
			domain.
				TokenRightParen {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, top)
	}
	return output, nil
}
func shouldPop(top,
	current domain.Token) bool {
	topInfo, topOk := domain.
		GetOpInfo(
			top.
				Value)
	curInfo, curOk :=

		domain.GetOpInfo(current.Value)
	if !topOk ||
		!curOk {
		return top.Type == domain.TokenFunction
	}
	if topInfo.Precedence >
		curInfo.Precedence {
		return true
	}
	if topInfo.
		Precedence == curInfo.
		Precedence &&
		!curInfo.
			IsRightAssociative {
		return true
	}
	return false
}
func isDigitOrDot(ch rune) bool {
	return unicode.
		IsDigit(ch) ||
		ch == '.'
}
func isAlphaUnderscore(ch rune) bool {
	return unicode.
		IsLetter(ch) || ch ==
		'_'
}
