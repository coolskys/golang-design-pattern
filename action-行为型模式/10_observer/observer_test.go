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
	subject.Notify("do your thing")

	// 移除接收对象
	subject.Detach(reader1)
	subject.Notify("do your thing again")

	// Output:
	// reader1 receive do your thing
	// reader2 receive do your thing
	// reader3 receive do your thing
	// writer1 write do your thing
	// reader2 receive do your thing again
	// reader3 receive do your thing again
	// writer1 write do your thing again
}

func TestAsyncEventBus(t *testing.T) {
	eventBus := NewAsyncEventBus()
	eventBus.Subscribe("topic1", sub1)
	eventBus.Subscribe("topic2", sub2)

	eventBus.Publish("topic1", "testA", "testB")

	eventBus.Publish("topic2", 1, 2, 3, 4, 5, 6)
}
