## 观察者模式

观察者模式用于触发联动。

一个对象的改变会触发其它观察者的相关动作，而此对象无需关心连动对象的具体实现。

观察者模式，就是一个对象发出一个操作，触发操作通知，所有观察者都会收到这个操作指令并进行自己的处理任务，输出自己的结果。

观察者，顾名思义就是观察某一个对象的行为，根据它的行为做出一些操作。


### 概念
观察者模式(Observer Pattern)，定义对象间的一种一对多的依赖关系，使得每当一个对象状态发生改变时，其相关依赖对象都得到通知并自动更新。观察者模式又叫做发布-订阅（Publish/Subscribe）模式、模型-视图（Model/View）模式、源-监听器（Source/Listener）模式或从属者（Dependents）模式。发布者对观察者唯一了解的是它实现了某个接口（观察者接口）。这种松散耦合的设计最大限度地减少了对象之间的相互依赖，因此使我们能够构建灵活的系统。

### 适用场景

1. 一个抽象模型有两个方面，其中一个方面依赖于另一个方面。将这些方面封装在独立的对象中使它们可以各自独立地改变和复用。
2. 一个对象的改变将导致其他一个或多个对象也发生改变，而不知道具体有多少对象将发生改变，可以降低对象之间的耦合度。
3. 一个对象必须通知其他对象，而并不知道这些对象是谁。
4. 需要在系统中创建一个触发链，A对象的行为将影响B对象，B对象的行为将影响C对象，可以使用观察者模式创建一种链式触发机制。

### 优点

1. 降低了目标与观察者之间的耦合关系，两者之间是抽象耦合关系。
2. 目标与观察者之间建立了一套触发机制。

### 缺点
1. 目标与观察者之间的依赖关系并没有完全解除，而且有可能出现循环引用。
2. 当观察者对象很多时，通知的发布会花费很多时间，影响程序的效率。



**参考文档**
[拒绝Go代码臃肿，其实在这几块可以用下观察者模式](https://mp.weixin.qq.com/s/4NqjkXVqFPamEc_QsyRipA)

