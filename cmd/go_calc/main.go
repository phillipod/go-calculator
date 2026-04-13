package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"go-calculator/internal/app"
	"go-calculator/internal/infix"
	"go-calculator/internal/rpn"
	"os"
	"strings"
)

func main() {
	calc := app.
		NewCalculatorService(infix.
			NewShuntingYard(), rpn.
			NewStackEvaluator())
	var rootCmd = &cobra.Command{Use: "go_calc",
		Short: "A scientific calculator with RPN engine"}
	var evalCmd = &cobra.Command{Use: "eval [expression]", Short: "Evaluate a mathematical expression", Args: cobra.MinimumNArgs(1), Run: func(cmd *cobra.Command, args []string) {
		expr := strings.Join(args, " ")
		result, err := calc.Calculate(expr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(app.FormatResult(result))
	}}
	var interactiveCmd = &cobra.Command{Use: "interactive", Short: "Start interactive calculator mode", Run: func(cmd *cobra.Command, args []string) {
		runInteractive(calc)
	}}
	var funcsCmd = &cobra.Command{Use: "funcs", Short: "List available functions", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available functions:")
		for _, fn := range app.AvailableFunctions() {
			fmt.Printf("  %s\n", fn)
		}
		fmt.Println("\nConstants:")
		for name, val := range app.AvailableConstants() {
			fmt.Printf("  %s = %g\n", name, val)
		}
	}}
	rootCmd.AddCommand(evalCmd, interactiveCmd, funcsCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
func runInteractive(calc *app.CalculatorService) {
	fmt.Println("go-calculator interactive mode")
	fmt.Println("Type 'exit' or 'quit' to quit, 'funcs' to list functions")
	fmt.Println()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.
			Text())
		if line == "" {
			continue
		}
		if line ==
			"exit" || line == "quit" {
			fmt.Println("Goodbye!")
			break
		}
		if line == "funcs" {
			fmt.Println("Functions:")
			for _, fn := range app.AvailableFunctions() {
				fmt.Printf("  %s(x)\n", fn)
			}
			fmt.Println("Constants: pi, e")
			continue
		}
		result,
			err := calc.Calculate(line)
		if err !=
			nil {
			fmt.
				Printf("Error: %v\n", err)
			continue
		}
		fmt.
			Printf("= %s\n", app.FormatResult(result))
	}
}
