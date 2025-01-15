package builder

import "fmt"

/*
	创建者模式

	该模式的主要目的是将对象的构建过程与其表示分离，使得同样的构建过程可以创建不同的表示

	将对象拆解成多个构建构成

	创建者模式通常由以下几个核心组件组成：

		产品（Product）：表示要创建的复杂对象。

		创建者（Builder）：定义了创建产品的各个部件的抽象方法。

		具体创建者（Concrete Builder）：实现了创建者接口，并负责具体产品的构建过程。

		导演（Director）：负责使用创建者来构建最终的产品。它并不直接与产品交互，而是通过调用创建者的方法来构建产品。
*/

//Builder 是生成器接口

type Builder interface {
	Part1()
	Part2()
	Part3()
	GetResult() interface{}
}

type Director struct {
	builder Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//Construct Product
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() interface{} {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() interface{} {
	return b.result
}

// 产品
type Host struct {
	name     string
	brand    string
	datetime string
}

func (c *Host) BuildHost(name string) {
	c.name = name
}

func (c *Host) BuildBrand(name string) {
	c.brand = name
}

func (c *Host) BuildDateTime(date string) {
	c.datetime = date
}

func (c *Host) GetResult() string {
	return fmt.Sprintf("host(%s-%s-%s)", c.name, c.brand, c.datetime)
}

type HostBuilder interface {
	BuildHost(string)
	BuildBrand(string)
	BuildDateTime(string)
	GetResult() string
}

type Screen struct {
	name  string
	brand string
}

func (c *Screen) BuildName(name string) {
	c.name = name
}

func (c *Screen) BuildBrand(name string) {
	c.brand = name
}

func (c *Screen) GetResult() string {
	return fmt.Sprintf("screen(%s-%s)", c.name, c.brand)
}

// 产品的构建方法,构建器，一系列构建方法组合

type ScreenBuilder interface {
	BuildName(string)
	BuildBrand(string)
	GetResult() string
}

// 产品构建过程

type ComputerDirector struct {
	hb HostBuilder
	sb ScreenBuilder
}

// 组装构建过程
func (c *ComputerDirector) Construct(name, brand string) string {
	c.hb.BuildHost(name)
	c.hb.BuildBrand(brand)

	c.sb.BuildName(name)
	c.sb.BuildBrand(brand)
	host := c.hb.GetResult()
	screen := c.sb.GetResult()

	return fmt.Sprintf("%s|%s", host, screen)
}

func NewComputerDirector(hbuilder HostBuilder, sbuilder ScreenBuilder) *ComputerDirector {
	return &ComputerDirector{
		hb: hbuilder,
		sb: sbuilder,
	}
}
