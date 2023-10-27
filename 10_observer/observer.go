package observer

import "fmt"

type Subject struct {
	observers []Observer
	context   string // 公共数据
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}

type Writer struct {
	name string
}

func NewWriter(name string) *Writer {
	return &Writer{
		name: name,
	}
}

func (w *Writer) Update(s *Subject) {
	fmt.Printf("%s write %s\n", w.name, s.context)
}
