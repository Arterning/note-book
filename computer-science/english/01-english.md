# doc

## spring

- https://docs.spring.io/spring-framework/docs/current/reference/html/index.html

- https://spring.io/guides/gs/scheduling-tasks/

# sentence

(in that the presence of setter methods and appropriate constructors on classes makes them easier to wire together in a test without having to set up service locator registries and similar structures).

（因为类上存在 setter 方法和适当的构造函数使它们更容易在测试中连接在一起，而无需设置服务定位器注册表和类似结构）。

# 词汇

- Appendix 附录
- advocates 主张
- integral
- Emphasizing 强调
- A consistent programming model 一致性的编程模型
- declarative transaction management
- Transaction bound event 事务绑定事件
- profound limitations 深刻的局限性
- Global transactions
- preffered 首选
- negatives 缺点 缺陷
- The negatives of EJB in general are so great that this is - not an attractive proposition
- EJB 总体上的缺点是如此之大，以至于这不是一个有吸引力的命题
- ensure correctness
- downside 缺点
- which can run over any underlying transaction
- infrastructure
- underlying 底层
- In particular 特别的
- notion 概念
- A transaction strategy is defined by a TransactionManager
- imperative 必要的事;祈使语气 重要紧急的事
- It is not tied to a lookup strategy 它与查找策略无关
- lookup 在这个上下文中是查找的意思 有的时候一个词汇多个意思 理解- 他的含义最好的方式是放在一个语句的上下文中理解
- This benefit alone makes Spring Framework transactions a worthwhile abstraction
- Worthwhile 值得的
- The salient point is that developers are not forced to do so.
- 最重要的一点是
- salient point 凸点；折点
- preceding 在前的;前面的
- This is due to the fact that JTA transactions are global transactions, which can enlist any transactional resource.
- 这是因为 JTA 事务是全局事务，可以调用任何事务资源。
- enlist 征集 ; 征募 ; 从军 ; 争取，谋取 ; 入伍
- focus purely on non-boilerplate persistence logic
- 只关注非样板持久性逻辑
- boilerplate
- invasive 侵入的
- It is not sufficient merely
- 仅仅这样是不够的
- merely 仅仅是
- explains the inner workings
- must run in the context of a transaction with read-write semantics
- 必须在具有读写语义的事务上下文中运行
- omit 省略
- resemble the following 类似于以下内容
- the Spring Framework’s transaction infrastructure code marks a transaction for rollback only in the case of runtime
- Spring 框架的事务基础结构代码仅在运行时标记要回滚的事务
- Spring Framework’s transaction infrastructure
- code marks 标记 code 还可以作为动词
- this process is quite invasive and tightly couples your code to the Spring Framework’s transaction infrastructure
- 这个过程是非常侵入性的，并且将你的代码与 Spring Framework 的事务基础设施紧密耦合。
- annotation-based approach
- Declaring transaction semantics directly in the Java source code puts the declarations much closer to the affected code
- 直接在 Java 源代码中声明事务语义会使声明更接近我们的代码
- semantics 语义
- undue coupling 过度耦合
- class-level annotation does not apply to ancestor classes up the class hierarchy
- 类级注释不适用于类层次结构中的祖先类
- Method visibility and @Transactional
  When you use transactional proxies with Spring’s standard configuration, you should apply the @Transactional annotation only to methods with public visibility. If you do annotate protected, private, or package-visible methods with the @Transactional annotation, no error is raised
- 方法可见性和@Transactional
  当您使用具有 Spring 标准配置的事务代理时，您应该仅将@Transactional 注释应用于具有公共可见性的方法。如果使用@Transactional 注释注释受保护、私有或包可见的方法，则不会出错。
- not exhibit 不显示
- merely 只是
- The Spring team recommends that you annotate only concrete classes
- The proxy-target-class attribute controls what type of transactional proxies are created for classes annotated with the @Transactional annotation. If proxy-target-class is set to true, class-based proxies are created. If proxy-target-class is false or if the attribute is omitted, standard JDK interface-based proxies are created
- 代理目标类属性控制为使用@Transactional 注释的类创建的事务代理类型。如果代理目标类设置为 true，则会创建基于类的代理。如果代理目标类为 false 或省略了该属性，则会创建基于 JDK 接口的标准代理
- precedence over 优先级高于
- The most derived location takes precedence when evaluating the transactional settings for a method.
- 在评估方法的事务设置时，派生最多的位置优先。
- 也就是方法级别的配置高于类级别的配置 这一点比较符合直觉 没什么好说的
- define custom composed annotations
- aforementioned 上述的
- Marshalling 编组
- Accommodate 容纳
- emphasis 强调
