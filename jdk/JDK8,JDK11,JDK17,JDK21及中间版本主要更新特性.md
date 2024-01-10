
When i use reacdt-native, it require JDK 17 , so i update to jdk 17...

but jdk 17 is a important version

let me see the jdk version changed..

https://pdai.tech/md/java/java8up/java12-17.html

---

#### JDK 17升级的必要性？

1. JDK 11 作为一个 LTS版本，它的商业支持时间框架比 JDK 8 短，**JDK 11 的 LTS 会提供技术支持直至 2023 年 9 月, 对应的补丁和安全警告等支持将持续至 2026 年**。JDK 17 作为下一代 LTS 将提供至少到 2026 年的支持时间框架；
2. Java系最为重要的开发框架**Spring Framework 6 和 Spring Boot 3对JDK版本的最低要求是JDK 17**；所以可以预见, 为了使用Spring最新框架，很多团队和开发者将被迫从Java 11（甚至Java 8)直接升级到Java 17版本。

#### [#](#jdk-11-升级到jdk-17-性能提升多少) JDK 11 升级到JDK 17 性能提升多少？

从规划调度引擎 OptaPlanner 项目（原文在[这里在新窗口打开](https://www.optaplanner.org/blog/2021/09/15/HowMuchFasterIsJava17.html)）对 JDK 17和 JDK 11 的性能基准测试进行了对比来看：

1. 对于 G1GC（默认），Java 17 比 Java 11 快 8.66%；
2. 对于 ParallelGC，Java 17 比 Java 11 快 6.54%；
3. Parallel GC 整体比 G1 GC 快 16.39%

简而言之，JDK17 更快，高吞吐量垃圾回收器比低延迟垃圾回收器更快。


---


Text block

```java
public static void main(String[] args) {
    String query = """
           SELECT * from USER \
           WHERE `id` = 1 \
           ORDER BY `id`, `name`;\
           """;
    System.out.println(query);
}

```

---

Record

```java
public record Person(String name, int age) {
    public static String address;

    public String getName() {
        return name;
    }
}
```

当用 Record 来声明一个类时，该类将自动拥有下面特征：

- 拥有一个构造方法
- 获取成员属性值的方法：name()、age()
- hashCode() 方法和 euqals() 方法
- toString() 方法
- 类对象和属性被 final 关键字修饰，不能被继承，类的示例属性也都被 final 修饰，不能再被赋值使用。
- 还可以在 Record 声明的类中定义静态属性、方法和示例方法。注意，不能在 Record 声明的类中定义示例字段，类也不能声明为抽象类等。

---

Sealed Interface

限制 Person类 只能被这三个类继承，不能被其他类继承，需要这么做。

```java
// 添加sealed修饰符，permits后面跟上只能被继承的子类名称
public sealed class Person permits Teacher, Worker, Student{ } //人

```