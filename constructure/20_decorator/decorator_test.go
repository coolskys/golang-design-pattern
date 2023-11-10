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

	var circle = &Circle{}
	var rectangle = &Rectangle{}

	var redshapeDecorator1 = RedShapeDecorator{
		ShapeDecorator: ShapeDecorator{
			decoratedShape: circle,
		},
	}

	var redshapeDecorator2 = RedShapeDecorator{
		ShapeDecorator: ShapeDecorator{
			decoratedShape: rectangle,
		},
	}
	redshapeDecorator1.draw()
	redshapeDecorator2.draw()


}
