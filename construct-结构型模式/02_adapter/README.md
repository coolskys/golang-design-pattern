# 适配器模式

适配器模式用于转换一种接口适配另一种接口。

实际使用中Adaptee一般为接口，并且使用工厂函数生成实例。

在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，又因为Go语言中非入侵式接口特征，其实Adapter也适配Adaptee接口。

* 使用现有组件：当你需要使用一个已经存在的类或接口，并且它的接口与你所期望的接口不兼容时，可以使用适配器模式。通过创建一个适配器，你可以将现有的类或接口适配成符合你需求的接口。

* 与第三方代码集成：当你需要与第三方库或框架进行集成，并且它们的接口与你的系统接口不匹配时，适配器模式也是一种解决方案。你可以通过创建适配器来使你的系统能够与第三方代码进行交互。

* 接口升级：当你的系统接口需要进行升级或改变时，但又不能直接修改现有接口的实现代码时，可以通过适配器模式来进行平滑过渡。你可以创建一个新的接口，并使用适配器将新接口适配到旧接口上。

* 统一多个接口：当你需要统一多个类或接口的行为，以提供统一的接口给客户端使用时，适配器模式可以帮助你实现这一目标。适配器可以将多个类或接口的方法进行适配和组合，形成一个统一的接口给客户端调用。

总之，适配器模式适用于需要将不兼容的接口进行适配的各种场景，帮助你实现接口之间的转换和集成。它能够提供一个灵活、可扩展和可维护的方式来解决不同接口之间的兼容性问题。
