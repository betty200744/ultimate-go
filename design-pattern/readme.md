
* 设计模式存在的价值是帮助程序员沟通，现在一般不用设计模式，现在都用第三方依赖， 语法糖


### Creational
- 简单工厂模式（Simple Factory）
- 工厂方法模式（Factory Method）
  - creates objects by a function without specifying the exact class to create
  - 通过一个function， 传入参数，创建不同的类，注意此时不需要指定特定的类
- 抽象工厂模式（Abstract Factory）
  - groups object factories that have a common theme
  - 相同类型的object一组里面即放在一个类里面， 之后用class method来获取具体的类
- 建造者模式（Builder）
  - constructs complex objects by separating construction and representation
  - 通过分步骤的方式来构造一个复杂的construction
- 原型模式（Prototype）
  - creates objects by cloning an existing object
  - 通过clone 方法来创建一个已经存在的object
- 单例模式（Singleton）
  - restricts object creation for a class to only one instance.
  - 创建一个单例， 无论new多少次， 确保值创建一次
### structual
- 适配器模式(Adapter)
- 桥接模式(Bridge)
- 组合模式(Composite)
- 装饰模式(Decorator)
- 外观模式(Facade)
- 享元模式(Flyweight)
- 代理模式(Proxy)

### behavioral
- 职责链模式(Chain of Responsibility) , 即middleware中间件的感觉
- 命令模式(Command)
- 解释器模式(Interpreter)
- 迭代器模式(Iterator)
- 中介者模式(Mediator)
- 备忘录模式(Memento)
- 观察者模式(Observer)
- 状态模式(State)
- 策略模式(Strategy), 登录用不同的账号登录， 即策略， 付款用不同的方式付款，也是策略， kafka消费数据用不同的策略
- 模板方法模式(Template Method)
- 访问者模式(Visitor)