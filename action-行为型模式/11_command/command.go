package command

import "fmt"

/*
命令模式本质是把某个对象的方法调用封装到对象中，方便传递、存储、调用。

示例中把主板单中的启动(start)方法和重启(reboot)方法封装为命令对象，再传递到主机(box)对象中。与两个按钮进行绑定：

* 第一个机箱(box1)设置按钮1(button1) 为开机按钮2(button2)为重启。
* 第二个机箱(box1)设置按钮2(button2) 为开机按钮1(button1)为重启。

从而得到配置灵活性。
*/
type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

type PlayCommand struct {
	mb *MotherBoard
}

func NewPlayCommand(mb *MotherBoard) *PlayCommand {
	return &PlayCommand{
		mb: mb,
	}
}

func (p *PlayCommand) Execute() {
	p.mb.PlaySound()
}

type CaculateCommand struct {
	mb *MotherBoard
}

func NewCaculateCommand(mb *MotherBoard) *CaculateCommand {
	return &CaculateCommand{
		mb: mb,
	}
}

func (c *CaculateCommand) Execute() {
	c.mb.Caculate()
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

func (*MotherBoard) PlaySound() {
	fmt.Printf("play sound\n")
}

func (*MotherBoard) Caculate() {
	fmt.Print("caculate\n")
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
