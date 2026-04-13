package domain

import "math"

type FuncInfo struct {
	Name  string
	Arity Arity
	Apply func(float64) float64
}

var functions = map[string]FuncInfo{"sqrt": {Name: "sqrt", Arity: Unary, Apply: math.Sqrt},
	"cbrt": {Name: "cbrt", Arity: Unary, Apply: math.Cbrt}, "abs": {Name: "abs", Arity: Unary, Apply: math.Abs}, "ceil": {Name: "ceil", Arity: Unary, Apply: math.Ceil}, "floor": {Name: "floor", Arity: Unary, Apply: math.Floor}, "round": {Name: "round", Arity: Unary, Apply: math.Round}, "exp": {Name: "exp", Arity: Unary, Apply: math.Exp}, "ln": {Name: "ln", Arity: Unary, Apply: math.Log}, "log": {Name: "log", Arity: Unary, Apply: math.Log10}, "log10": {Name: "log10", Arity: Unary, Apply: math.Log10}, "log2": {Name: "log2", Arity: Unary, Apply: math.Log2}, "sin": {Name: "sin", Arity: Unary, Apply: math.Sin}, "cos": {Name: "cos", Arity: Unary, Apply: math.Cos}, "tan": {Name: "tan", Arity: Unary, Apply: math.Tan}, "asin": {Name: "asin", Arity: Unary, Apply: math.Asin}, "acos": {Name: "acos", Arity: Unary, Apply: math.Acos}, "atan": {Name: "atan", Arity: Unary, Apply: math.Atan}, "sinh": {Name: "sinh", Arity: Unary, Apply: math.Sinh}, "cosh": {Name: "cosh", Arity: Unary, Apply: math.Cosh}, "tanh": {Name: "tanh", Arity: Unary, Apply: math.Tanh}, "fact": {Name: "fact", Arity: Unary, Apply: factorial}, "deg": {Name: "deg", Arity: Unary, Apply: func(x float64) float64 {
		return x * 180 / math.Pi
	}}, "rad": {Name: "rad", Arity: Unary, Apply: func(x float64) float64 {
		return x * math.Pi / 180
	}}}
var Constants = map[string]float64{"pi": math.Pi,
	"e": math.E,
}

func factorial(
	x float64) float64 {
	if x < 0 || x !=
		math.Trunc(x) {
		return math.NaN()
	}
	result := 1.0
	for i := 2.0; i <= x; i++ {
		result *= i
	}
	return result
}
func GetFuncInfo(name string) (FuncInfo, bool) {
	info,
		ok := functions[name]
	return info, ok
}
func IsFunction(
	name string) bool {
	_,
		ok := functions[name]
	return ok
}
func IsConstant(
	name string) bool {
	_,
		ok := Constants[name]
	return ok
}
func GetConstant(name string) (float64,
	bool) {
	val,
		ok := Constants[name]
	return val, ok
}
