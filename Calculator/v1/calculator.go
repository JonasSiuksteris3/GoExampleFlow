package v1

import (
	"fmt"
)

func addNumbers(a, b int) int {
	return a + b
}

func main() {
	result := addNumbers(5, 3)
	fmt.Println("Result:", result)
}
