# 职责链模式

职责链模式用于分离不同职责，并且动态组合相关职责。

Golang实现职责链模式时候，因为没有继承的支持，使用链对象包涵职责的方式，即：

* 链对象包含当前职责对象以及下一个职责链。
* 职责对象提供接口表示是否能处理对应请求。
* 职责对象提供处理函数处理相关职责。

同时可在职责链类中实现职责接口相关函数，使职责链对象可以当做一般职责对象使用。

一般可用于对象路由


