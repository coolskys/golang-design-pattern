package memento

import (
	"testing"
)

func TestExampleGame(t *testing.T) {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.Status()
	progress := game.Save() // 保存到memento对象

	game.Play(-2, -3)
	game.Status()

	game.Load(progress)
	game.Status()

	// Output:
	// Current HP:10, MP:10
	// Current HP:7, MP:8
	// Current HP:10, MP:10
}
