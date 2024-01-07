package main

import (
	"fmt"
)

// Operation interface for different arithmetic operations
type Operation interface {
	Execute(a, b int) (int, error)
}

// AddOperation struct for addition
type AddOperation struct{}

func (op AddOperation) Execute(a, b int) (int, error) {
	return a + b, nil
}

// SubtractOperation struct for subtraction
type SubtractOperation struct{}

func (op SubtractOperation) Execute(a, b int) (int, error) {
	return a - b, nil
}

// MultiplyOperation struct for multiplication
type MultiplyOperation struct{}

func (op MultiplyOperation) Execute(a, b int) (int, error) {
	return a * b, nil
}

// AdvancedCalculator struct that uses Operation interface
type AdvancedCalculator struct {
	operation Operation
}

func (ac *AdvancedCalculator) SetOperation(op Operation) {
	ac.operation = op
}

func (ac *AdvancedCalculator) Calculate(a, b int) (int, error) {
	if ac.operation == nil {
		return 0, fmt.Errorf("no operation set")
	}
	return ac.operation.Execute(a, b)
}

func main() {
	calc := AdvancedCalculator{}

	// Addition
	calc.SetOperation(AddOperation{})
	result, err := calc.Calculate(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Addition Result:", result)
	}

	// Subtraction
	calc.SetOperation(SubtractOperation{})
	result, err = calc.Calculate(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Subtraction Result:", result)
	}

	// Multiplication
	calc.SetOperation(MultiplyOperation{})
	result, err = calc.Calculate(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Multiplication Result:", result)
	}
}
