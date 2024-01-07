package v2

import (
	"fmt"
)

type Calculator struct {
	total int
}

func (c *Calculator) add(a, b int) {
	c.total = a + b
}

func (c *Calculator) displayTotal() {
	fmt.Println("Total:", c.total)
}

func main() {
	calc := Calculator{}
	calc.add(5, 3)
	calc.displayTotal()
}
