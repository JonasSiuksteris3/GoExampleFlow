package v3

import (
	"errors"
	"fmt"
)

type AdvancedCalculator struct {
	total     int
	operation string
}

func (ac *AdvancedCalculator) add(a, b int) {
	ac.total = a + b
	ac.operation = "add"
}

func (ac *AdvancedCalculator) subtract(a, b int) {
	ac.total = a - b
	ac.operation = "subtract"
}

func (ac *AdvancedCalculator) multiply(a, b int) {
	ac.total = a * b
	ac.operation = "multiply"
}

func (ac *AdvancedCalculator) displayResult() error {
	if ac.operation == "" {
		return errors.New("no operation performed")
	}
	fmt.Printf("Result after %s: %d\n", ac.operation, ac.total)
	return nil
}

func main() {
	calc := AdvancedCalculator{}
	calc.add(10, 5)
	err := calc.displayResult()
	if err != nil {
		fmt.Println("Error:", err)
	}
	calc.multiply(3, 3)
	err = calc.displayResult()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
