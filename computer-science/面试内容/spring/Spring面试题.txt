什么是依赖注入
就是你不需要创建对象 只需要描述如何创建对象 具体的创建工作由容器去做

BeanFactory->ApplicationContext->WebApplicationContext

SpringIOC的原理就是工厂模式和反射机制

Spring自动装备
byName
byType

AOP 有哪些实现方式？
动态代理 运行时增强
编译时增强
类加载时增强


静态代理
class Proxy {
  realObject;

  func() {
    ...
    ..
    realObject.doSomething();
    ...
    ...

  }

}


DispatcherServlet 的工作流程？





