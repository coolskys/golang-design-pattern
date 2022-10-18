package decorator

import (
	"fmt"
	"testing"
)

func TestExampleDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WrapMulDecorator(c, 8)
	c = WrapAddDecorator(c, 10)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
	// Output:
	// res 80
}
