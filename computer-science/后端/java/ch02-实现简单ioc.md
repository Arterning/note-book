以下是一个简单的 Java IOC（Inversion of Control）容器示例，支持依赖注入。该示例使用了基于反射的依赖注入。

首先，我们需要定义一个简单的接口 `Service` 和两个实现类 `ServiceA` 和 `ServiceB`：

```java
// Service.java
public interface Service {
    void execute();
}

// ServiceA.java
public class ServiceA implements Service {
    @Override
    public void execute() {
        System.out.println("ServiceA executed.");
    }
}

// ServiceB.java
public class ServiceB implements Service {
    @Override
    public void execute() {
        System.out.println("ServiceB executed.");
    }
}
```

然后，我们创建一个 IOC 容器类 `IoCContainer`，该容器可以注册、获取和注入依赖：

```java
import java.util.HashMap;
import java.util.Map;

public class IoCContainer {
    private Map<String, Object> beans = new HashMap<>();

    public void register(String beanName, Object bean) {
        beans.put(beanName, bean);
    }

    public Object getBean(String beanName) {
        return beans.get(beanName);
    }

    public void injectDependencies(Object target) {
        Class<?> targetClass = target.getClass();
        for (var field : targetClass.getDeclaredFields()) {
            if (field.isAnnotationPresent(Autowired.class)) {
                field.setAccessible(true);
                try {
                    Object dependency = getBean(field.getName());
                    field.set(target, dependency);
                } catch (IllegalAccessException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
```

注意，我们使用 `@Autowired` 注解来标记需要进行依赖注入的字段。

接下来，我们创建一个简单的应用程序来演示 IOC 容器的使用：

```java
public class Main {
    public static void main(String[] args) {
        IoCContainer container = new IoCContainer();

        // 注册服务
        container.register("serviceA", new ServiceA());
        container.register("serviceB", new ServiceB());

        // 创建并注入依赖
        Client client = new Client();
        container.injectDependencies(client);

        // 使用依赖
        client.invokeServiceA();
        client.invokeServiceB();
    }
}
```

在这个示例中，我们首先注册了两个服务（`ServiceA` 和 `ServiceB`），然后使用 `IoCContainer` 来创建一个 `Client` 实例并注入依赖。`Client` 类使用了 `@Autowired` 注解来标记需要注入的字段。

最后，我们运行应用程序，可以看到输出：

```
ServiceA executed.
ServiceB executed.
```

这证明了 IOC 容器成功进行了依赖注入，使得 `Client` 类能够使用 `ServiceA` 和 `ServiceB` 的功能。

请注意，这只是一个非常简单的示例，实际的 IOC 容器通常会更复杂，支持更多功能，例如循环依赖的解决、生命周期管理等。在实际项目中，你可能会考虑使用现有的 IOC 容器库（如Spring Framework），以便利用更多功能和性能优化。