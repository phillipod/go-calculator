# go-calculator

A command-line scientific calculator written in Go. It parses infix expressions using the shunting-yard algorithm, converts them to Reverse Polish Notation (RPN), and evaluates them with a stack-based engine.

## Install

```bash
go build -o go_calc ./cmd/go_calc
```

## Usage

### Evaluate an expression

```bash
go_calc eval "2 + 3 * 4"         # 14
go_calc eval "sqrt(16) + 2^3"     # 12
go_calc eval "sin(pi / 2)"        # 1
go_calc eval "fact(5)"            # 120
```

### Interactive mode

```bash
go_calc interactive
```

Opens a REPL where you can type expressions. Type `funcs` to list available functions, `exit` or `quit` to leave.

### List available functions

```bash
go_calc funcs
```

## Supported Operations

| Operator | Description       | Example        |
|----------|-------------------|----------------|
| `+`      | Addition          | `2 + 3`        |
| `-`      | Subtraction       | `5 - 2`        |
| `*`      | Multiplication    | `3 * 4`        |
| `/`      | Division          | `10 / 2`       |
| `%`      | Modulo            | `10 % 3`       |
| `^`      | Exponentiation    | `2 ^ 10`       |

Unary minus is supported: `-5 + 3` evaluates to `-2`.

## Functions

`sqrt`, `cbrt`, `abs`, `ceil`, `floor`, `round`, `exp`, `ln`, `log`, `log10`, `log2`, `sin`, `cos`, `tan`, `asin`, `acos`, `atan`, `sinh`, `cosh`, `tanh`, `fact`, `deg`, `rad`

## Constants

`pi` (3.14159...) and `e` (2.71828...)

## Project Structure

```
cmd/go_calc/          CLI entry point (Cobra commands)
internal/
  app/                Calculator service (orchestrates translation + evaluation)
  domain/             Core types: Token, OpInfo, FuncInfo, operations, functions
  port/               Interfaces: Translator, Evaluator, Calculator
  infix/              Shunting-yard parser (infix -> RPN tokens)
  rpn/                Stack-based RPN evaluator
```

## Dependencies

- [cobra](https://github.com/spf13/cobra) - CLI framework

## License

GPL-2.0
