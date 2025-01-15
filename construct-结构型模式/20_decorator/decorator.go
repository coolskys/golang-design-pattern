package decorator

import "fmt"

/*
装饰模式使用对象组合的方式动态改变或增加对象行为。

Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。

使用匿名组合，在装饰器中不必显式定义转调原对象方法。
*/

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

// 抽象组件
type Shape interface {
	draw()
}

// 具体组件
type Circle struct{}

func (c *Circle) draw() {
	fmt.Println("画一个圆形")
}

type Rectangle struct{}

func (r *Rectangle) draw() {
	fmt.Println("画一个矩形")
}

// 抽象装饰器
type ShapeDecorator struct {
	decoratedShape Shape
}

func (sd *ShapeDecorator) draw() {
	sd.decoratedShape.draw()
}

// 具体装饰器
type RedShapeDecorator struct {
	ShapeDecorator
}

func (rsd *RedShapeDecorator) draw() {
	rsd.decoratedShape.draw()
	fmt.Println("填充红色")
}
