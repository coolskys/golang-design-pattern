package observer

import "testing"

func TestExampleObserver(t *testing.T) {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	writer1 := NewWriter("writer1")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)
	subject.Attach(writer1)
	// 通知触发所有实例对象更新
	subject.UpdateContext("do your thing")
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}
