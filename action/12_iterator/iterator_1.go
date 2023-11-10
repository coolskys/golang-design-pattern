package iterator

// Iterator 迭代器接口
type Iterator1 interface {
	HasNext() bool
	Next() interface{}
}

// ConcreteIterator 具体迭代器类型
type ConcreteIterator struct {
	data     []interface{} // 聚合对象的数据
	position int           // 当前位置
}

// NewConcreteIterator 创建一个具体迭代器实例
func NewConcreteIterator(data []interface{}) *ConcreteIterator {
	return &ConcreteIterator{
		data:     data,
		position: 0,
	}
}

// HasNext 判断是否还有下一个元素
func (it *ConcreteIterator) HasNext() bool {
	return it.position < len(it.data)
}

// Next 返回下一个元素
func (it *ConcreteIterator) Next() interface{} {
	if it.HasNext() {
		value := it.data[it.position]
		it.position++
		return value
	}
	return nil
}

// Aggregate 聚合对象接口
type Aggregate1 interface {
	CreateIterator() Iterator1
}

// ConcreteAggregate 具体聚合对象类型
type ConcreteAggregate struct {
	data []interface{} // 聚合对象的数据
}

// CreateIterator 创建一个迭代器
func (a *ConcreteAggregate) CreateIterator() Iterator1 {
	return NewConcreteIterator(a.data)
}
