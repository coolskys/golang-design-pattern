package builder

import (
	"testing"
)

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123" {
		t.Fatalf("Builder1 fail expect 123 acture %s", res)
	}
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != 6 {
		t.Fatalf("Builder2 fail expect 6 acture %d", res)
	}
}

func TestBuildComputer(t *testing.T) {
	host := &Host{}
	screen := &Screen{}
	director := NewComputerDirector(host, screen)
	res := director.Construct("kewu", "xiaomi")

	if res != "host(kewu-xiaomi-)|screen(kewu-xiaomi)" {
		t.Fatalf("builder computer acture: %s", res)
	}
}
