package example

import (
	"fmt"
	"testing"
	"time"
)

func TestTrafficLight(t *testing.T) {
	var tl = NewTrafficLight(nil)
	fmt.Printf("%s: %s\n", time.Now().Format(time.DateTime), tl.State.State())
	var ticker = time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("%s: %s\n", time.Now().Format(time.DateTime), tl.State.Next(tl))
		}
	}
}
