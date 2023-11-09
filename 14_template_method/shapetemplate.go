package templatemethod

import "fmt"

// 图形接口
type Shapes interface {
	Draw()
}

type shapeTemplate struct {
	Draws
	color string
}

func newShapeTemplate(drawer Draws, color string) *shapeTemplate {
	return &shapeTemplate{
		Draws: drawer,
		color: color,
	}
}

// 模板流程
func (t *shapeTemplate) Draw() {
	// 调用画笔画形状
	t.Draws.draw()
	// 填充颜色
	fmt.Println("填充", t.color)
}

// 定义一个子类的接口用于画出形状
type Draws interface {
	draw()
}

// 画笔拥有模板方法，
type CircleDrawer struct {
	*shapeTemplate
}

func NewCircleDrawer(color string) Shapes {
	circleDrawer := &CircleDrawer{}
	t := newShapeTemplate(circleDrawer, color)
	circleDrawer.shapeTemplate = t
	return circleDrawer
}

// 子类实现画形状的方法
func (c *CircleDrawer) draw() {
	fmt.Println("画一个圆")
}

type RectangleDrawer struct {
	*shapeTemplate
}

func NewRectangleDrawer(color string) Shapes {
	rectangle := &RectangleDrawer{}
	t := newShapeTemplate(rectangle, color)

	rectangle.shapeTemplate = t
	return rectangle
}

func (c *RectangleDrawer) draw() {
	fmt.Println("画一个矩形")
}
