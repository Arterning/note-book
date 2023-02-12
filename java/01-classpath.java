/**
 * 在Java中，我们经常听到classpath这个东西。网上有很多关于“如何设置classpath”的文章，但大部分设置都不靠谱。

到底什么是classpath？

classpath是JVM用到的一个环境变量，它用来指示JVM如何搜索class。

因为Java是编译型语言，源码文件是.java，而编译后的.class文件才是真正可以被JVM执行的字节码。因此，JVM需要知道，如果要加载一个abc.xyz.Hello的类，应该去哪搜索对应的Hello.class文件。

所以，classpath就是一组目录的集合，它设置的搜索路径与操作系统相关。例如，在Windows系统上，用;分隔，带空格的目录用""括起来，可能长这样：

C:\work\project1\bin;C:\shared;"D:\My Documents\project1\bin"

在Linux系统上，用:分隔，可能长这样：

/usr/shared:/usr/local/bin:/home/liaoxuefeng/bin

现在我们假设classpath是.;C:\work\project1\bin;C:\shared，当JVM在加载abc.xyz.Hello这个类时，会依次查找：

<当前目录>\abc\xyz\Hello.class
C:\work\project1\bin\abc\xyz\Hello.class
C:\shared\abc\xyz\Hello.class



注意到.代表当前目录。如果JVM在某个路径下找到了对应的class文件，就不再往后继续搜索。如果所有路径下都没有找到，就报错。

classpath的设定方法有两种：

在系统环境变量中设置classpath环境变量，不推荐；

在启动JVM时设置classpath变量，推荐。

我们强烈不推荐在系统环境变量中设置classpath，那样会污染整个系统环境。在启动JVM时设置classpath才是推荐的做法。实际上就是给java命令传入-classpath或-cp参数：

java -classpath .;C:\work\project1\bin;C:\shared abc.xyz.Hello
或者使用-cp的简写：

java -cp .;C:\work\project1\bin;C:\shared abc.xyz.Hello
没有设置系统环境变量，也没有传入-cp参数，那么JVM默认的classpath为.，即当前目录：



java abc.xyz.Hello
上述命令告诉JVM只在当前目录搜索Hello.class。

在IDE中运行Java程序，IDE自动传入的-cp参数是当前工程的bin目录和引入的jar包。

通常，我们在自己编写的class中，会引用Java核心库的class，例如，String、ArrayList等。这些class应该上哪去找？

有很多“如何设置classpath”的文章会告诉你把JVM自带的rt.jar放入classpath，但事实上，根本不需要告诉JVM如何去Java核心库查找class，JVM怎么可能笨到连自己的核心库在哪都不知道？


更好的做法是，不要设置classpath！默认的当前目录.对于绝大多数情况都够用了。

假设我们有一个编译后的Hello.class，它的包名是com.example，当前目录是C:\work，那么，目录结构必须如下：

C:\work
└─ com
   └─ example
      └─ Hello.class
运行这个Hello.class必须在当前目录下使用如下命令：

C:\work> java -cp . com.example.Hello
JVM根据classpath设置的.在当前目录下查找com.example.Hello，即实际搜索文件必须位于com/example/Hello.class。如果指定的.class文件不存在，或者目录结构和包名对不上，均会报错。
 */