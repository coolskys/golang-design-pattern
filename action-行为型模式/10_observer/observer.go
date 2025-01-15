package observer

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type ISubject interface {
	Attach(o Observer)
	Detach(o Observer)
	Notify(msg string)
}

type Subject struct {
	observers []Observer
	msg       string
}

func NewSubject() ISubject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Detach(o Observer) {
	for i, o := range s.observers {
		if o == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) Notify(msg string) {
	s.msg = msg
	s.notify()
}

// 通过此接口操作所有订阅者对象
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
	fmt.Printf("%s receive %s\n", r.name, s.msg)
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
	fmt.Printf("%s write %s\n", w.name, s.msg)
}

// 事件总线实现
type EventBus interface {
	Subscribe(topic string, f interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: make(map[string][]reflect.Value, 0),
		lock:     sync.Mutex{},
	}
}

func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	if f == nil {
		return errors.New("f is nil")
	}

	val := reflect.ValueOf(f)
	if val.Kind() != reflect.Func {
		return errors.New("f is not a function")
	}

	handler, ok := bus.handlers[topic]
	if ok {
		handler = append(handler, val)
	} else {
		handler = []reflect.Value{val}
	}
	bus.handlers[topic] = handler
	return nil
}

// 异步
func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	if bus.handlers == nil {
		return
	}

	h, ok := bus.handlers[topic]
	if !ok {
		return
	}

	var params []reflect.Value
	for _, val := range args {
		v := reflect.ValueOf(val)
		params = append(params, v)
	}
	for _, hs := range h {
		go hs.Call(params)
	}
}

func sub1(msg1, msg2 string) {
	fmt.Printf("sub1: %s, %s do something!\n", msg1, msg2)
}

func sub2(args ...interface{}) {
	fmt.Printf("sub2: %v, do something!\n", args)
}
