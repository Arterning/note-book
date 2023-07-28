# jetbrain产品激活

# 配置
-javaagent:D:/hack/ja-netfilter.jar


# 什么是javaagent？

javaagent是java命令的一个参数，参数 javaagent可以用于指定一个 jar 包，并且对该 java 包有2个要求：

这个 jar 包的 MANIFEST.MF 文件必须指定 Premain-Class 项。

Premain-Class 指定的那个类必须实现 premain() 方法。

premain() 方法，从字面上理解，就是运行在 main() 函数之前的的类。当Java 虚拟机启动时，在执行 main() 函数之前，JVM 会先运行-javaagent所指定 jar 包内 Premain-Class 这个类的 premain() 方法 。

# ja-netfilter是什么，它做了哪些事？
从它的名字中就可以看出来，它是一个网络过滤器，可以屏蔽指定规则的网络请求，所以说大家将其理解为一个网络阻断器、网络拦截器更加合适。

配合上面的javaagent概念来理解，将其配置到idea中，它将在idea启动的时候拦截掉janf_config.txt文件中指定的网络请求。

在整个激活过程中，它的作用可以看做是欺骗Jetbrains官方，让他永远无法成功的知道你当前使用的激活码到底是不是过期了。


# mymap是什么，它有什么作用，不要它会怎么样？
我赌你们这些看文章的人都是没有付费购买过正版激活码的，那么请回忆一下你之前在网上找的一些激活码，填进去之后，在idea中显示的是不是Licensed to xxx的固定信息，若是你不想显示他们的这些信息，你想显示你自己的信息，该怎么办呢？此时mymap就能帮到你了，借助于它，你可以灵活自定义这些LicenseName等一些激活信息。

所以说，应该将ja-netfilter和mymap看做是两个单独的个体，只不过ja-netfilter可以独立工作，而mymap则是需要依托于ja-netfilter来发挥它的作用，在ja-netfilter配置正确的情况下，它已经破坏了idea的校验机制，无需mymap就能达到让你长期使用idea的效果了。

这么解释应该能明白了，如果你没有使用mymap，那也不会影响你激活idea，因为LicenseName和激活时长之类的信息自定义与否，关系不大，它们在我看来只是一个障眼法，欺骗你自己罢了，让你看着爽，其实也还是镜花水月。


