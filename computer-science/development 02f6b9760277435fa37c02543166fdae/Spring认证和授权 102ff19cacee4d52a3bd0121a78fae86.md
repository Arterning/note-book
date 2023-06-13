# Spring认证和授权

Created time: May 14, 2023 8:11 PM

总之我们只要知道这玩意是用在spring framwork 里保证程序安全的就行了。安全这个领域涵盖非常的广泛且深奥，二狗也没能力触碰太深。这里我们只涉及Authentication和Authorization两个概念

- 认证（Authentication）

解决你是谁的问题，具体表现为注册与登录

- 授权（Authorization）

解决你能干什么的问题，你登录后有哪些权限。

现在让你写一个登录功能（Authentication）你怎么写？很自然的思路是不是把用户提交的信息和我们保存的信息做个比较，如果对上了就登录成功。其实spring security整体也是这样的，只是流程化后，兼顾扩展导致搞的很复杂。

Spring Security的整体原理为：

1. 当http请求进来时，使用severlet的Filter来拦截。
2. 提取http请求中的认证信息，例如username和password，或者Token。
3. 从数据库（或者其他地方，例如Redis）中查询用户注册时的信息，然后进行比对，相同则认证成功，反之失败。

主体就是这么简单，然后只有抓住这个主体思路才不容易被Spring Security绕晕...

下图展示了Spring Security的一些Filter，其中`UsernamePasswordAuthenticationFilter`很重要，它是Authentication的开始。

```jsx
<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-security</artifactId>
    </dependency>
</dependencies>
```

在配置文件中配置用户名和密码

```jsx
spring:
  security:
    user:
      name: shusheng007
      password: ss007
```