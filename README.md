# Go 语言设计模式

[![Build Status](https://travis-ci.org/senghoo/golang-design-pattern.svg?branch=master)](https://travis-ci.org/senghoo/golang-design-pattern)

Go 语言设计模式的实例代码

## 创建型模式

* [简单工厂模式（Simple Factory）](https://github.com/senghoo/golang-design-pattern/tree/master/00_simple_factory)
* [单例模式（Singleton）](https://github.com/senghoo/golang-design-pattern/tree/master/03_singleton)
* [工厂方法模式（Factory Method）](https://github.com/senghoo/golang-design-pattern/tree/master/04_factory_method)
* [抽象工厂模式（Abstract Factory）](https://github.com/senghoo/golang-design-pattern/tree/master/05_abstract_factory)
* [创建者模式（Builder）](https://github.com/senghoo/golang-design-pattern/tree/master/06_builder)
* [原型模式（Prototype）](https://github.com/senghoo/golang-design-pattern/tree/master/07_prototype)

## 结构型模式

* [外观模式（Facade）](https://github.com/senghoo/golang-design-pattern/tree/master/01_facade)
* [适配器模式（Adapter）](https://github.com/senghoo/golang-design-pattern/tree/master/02_adapter)
* [代理模式（Proxy）](https://github.com/senghoo/golang-design-pattern/tree/master/09_proxy)
* [组合模式（Composite）](https://github.com/senghoo/golang-design-pattern/tree/master/13_composite)
* [享元模式（Flyweight）](https://github.com/senghoo/golang-design-pattern/tree/master/18_flyweight)
* [装饰模式（Decorator）](https://github.com/senghoo/golang-design-pattern/tree/master/20_decorator)
* [桥模式（Bridge）](https://github.com/senghoo/golang-design-pattern/tree/master/22_bridge)

## 行为型模式

* [中介者模式（Mediator）](https://github.com/senghoo/golang-design-pattern/tree/master/08_mediator)
* [观察者模式（Observer）](https://github.com/senghoo/golang-design-pattern/tree/master/10_observer)
* [命令模式（Command）](https://github.com/senghoo/golang-design-pattern/tree/master/11_command)
* [迭代器模式（Iterator）](https://github.com/senghoo/golang-design-pattern/tree/master/12_iterator)
* [模板方法模式（Template Method）](https://github.com/senghoo/golang-design-pattern/tree/master/14_template_method)
* [策略模式（Strategy）](https://github.com/senghoo/golang-design-pattern/tree/master/15_strategy)
* [状态模式（State）](https://github.com/senghoo/golang-design-pattern/tree/master/16_state)
* [备忘录模式（Memento）](https://github.com/senghoo/golang-design-pattern/tree/master/17_memento)
* [解释器模式（Interpreter）](https://github.com/senghoo/golang-design-pattern/tree/master/19_interpreter)
* [职责链模式（Chain of Responsibility）](https://github.com/senghoo/golang-design-pattern/tree/master/21_chain_of_responsibility)
* [访问者模式（Visitor）](https://github.com/senghoo/golang-design-pattern/tree/master/23_visitor)

## 模式使用口诀

单例模式（Singleton）：唯一实例，全局存储。
工厂模式（Factory）：封装创建，统一交付。
抽象工厂模式（Abstract Factory）：一族产品，共同创造。
建造者模式（Builder）：步骤分明，构建对象。
原型模式（Prototype）：克隆复制，快速生成。
适配器模式（Adapter）：接口转换，协作无碍。
桥接模式（Bridge）：抽象实现，解耦衔接。
装饰器模式（Decorator）：增强功能，动态包裹。
组合模式（Composite）：整体部分，统一对待。
外观模式（Facade）：简化接口，易于使用。
享元模式（Flyweight）：共享资源，节约开销。
代理模式（Proxy）：控制访问，安全可靠。
策略模式（Strategy）：算法替换，灵活应对。
模板方法模式（Template Method）：骨架固定，细节自由。
观察者模式（Observer）：消息订阅，及时通知。
迭代器模式（Iterator）：遍历集合，统一访问。
职责链模式（Chain of Responsibility）：责任传递，按需处理。
命令模式（Command）：行为封装，可撤销执行。
备忘录模式（Memento）：状态保存，回溯可行。
状态模式（State）：状态切换，行为变化。
访问者模式（Visitor）：双重分派，动态操作。
中介者模式（Mediator）：解耦合，统一调度。
解释器模式（Interpreter）：语法解析，执行操作。
亨元模式（Co享元模式）：共享细粒，节约存储。