package command

import (
	"fmt"
	"testing"
)

func TestExampleCommand(t *testing.T) {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)
	fmt.Println("box1 exec command")
	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButton1()
	box1.PressButton2()

	fmt.Println("box2 exec command")
	box2 := NewBox(rebootCommand, startCommand)
	box2.PressButton1()
	box2.PressButton2()
	// Output:
	// system starting
	// system rebooting
	// system rebooting
	// system starting
}
