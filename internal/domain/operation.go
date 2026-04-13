package domain

type Arity int

const (
	Unary  Arity = 1
	Binary Arity = 2
)

type OpInfo struct {
	Symbol             string
	Precedence         int
	IsRightAssociative bool
	Arity              Arity
}

var operations = map[string]OpInfo{"+": {Symbol: "+", Precedence: 1, IsRightAssociative: false, Arity: Binary}, "-": {Symbol: "-",
	Precedence: 1, IsRightAssociative: false, Arity: Binary}, "*": {Symbol: "*", Precedence: 2, IsRightAssociative: false, Arity: Binary}, "/": {Symbol: "/", Precedence: 2, IsRightAssociative: false, Arity: Binary}, "%": {Symbol: "%", Precedence: 2, IsRightAssociative: false, Arity: Binary}, "^": {Symbol: "^", Precedence: 3, IsRightAssociative: true, Arity: Binary}, "neg": {Symbol: "neg", Precedence: 4, IsRightAssociative: false, Arity: Unary}}

func GetOpInfo(op string) (OpInfo, bool) {
	info, ok := operations[op]
	return info,
		ok
}
func IsOperator(
	op string) bool {
	_, ok :=
		operations[op]
	return ok
}
