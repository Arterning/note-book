
/**
go导入的是路径 
比如
import "testmodule/foo"
导入时，是按照目录导入。导入目录后，可以使用这个目录下的所有包。
出于习惯，包名和目录名通常会设置成一样，所以会让你有一种你导入的是包的错觉。
所以go语言中不需要写大量的import
因为一个import就导入了目录下所有的包 这就有很多东西了

gopath是存放依赖的位置，具体是gopath/src
go mod会在build的时候在gopath中寻找你依赖的包
go 相比 java 确实相对简单不少 java classpath有很多个
所以go会把整个gopath下面的包都找一遍，如果找不到你要
依赖的包，就会报错。


横向对比一下java
你如果需要引入jar包的所有类 需要使用*符号 而go是不需要的
相当于go默认导入所有东西 默认加了*符号 导入这个目录下所有的东西
需要import a.b.c.*
java中的jar包无非就是把依赖的内容打包成zip格式而已
他的本质其实就是一个目录

maven的详细原理其实并没有详细思考？
这里不妨猜一下？
因为java和go还是有区别的 java的classpath有很多个
实际上一个dependency依赖就相当于一个classpath 一个jar包本质就是一个目录
此外在maven项目中 默认认为src/main/resources 和 src/main/java 目录也是classpath 因此如果你的依赖或配置文件在这些目录中，那么它们将被自动加载。

java当初设计这么多classpath的目的也许是为了灵活性
这样的好处应该是查找依赖的速度比较快 因为classpath已经限定了
你只能去哪些目录中寻找依赖。而不是要去整个maven仓库找一遍

**/

/**

后记

模块化是语言的一个重要特性 很大程度上决定和体现了编程语言的工程能力
通过比较不同编程语言的模块化解决方案 可以让我们对这一问题有更深入的思考
**/