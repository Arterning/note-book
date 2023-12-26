## rust核心特性是所有权
rust没有垃圾搜集器
在编译的阶段就做好了内存管理的工作

Stack 栈内存
存储在Stack上的数据必须是已知的固定的大小
分配非常快
当函数被调用的时候
函数的本地变量会被压入Stack 函数调用结束后弹出这些变量

Heap 堆内存
存储在Heap上的数据的内存空间是未知的




## 所有权解决的问题
1. 跟踪代码的哪些部分正在使用heap的哪些数据
2. 最小化heap的重复数据量
3. 清理heap上的未使用的数据

String类型是存储在Heap上 因为他是可变的
字符串字面值是硬编码写入可执行文件的

当拥有某个值的变量走出他的作用域以后，就自动释放内存





## 移动
let x =5;
let y = x;
x y都是有效的 因为x是整数类型的 对于这种类型的值的复制操作是非常快速的
整数是已知大小的值 所有两个5都被压入stack中


## 所有权转移
let s1=String::From("hello")
let s2=s1

s1被废弃 无法再被使用

rust不会自动创建对象的深拷贝 如果需要对heap上的对象进行deep clone 需要使用clone()

## 克隆
let s1=String::From("hello")
let s2=s1.clone()
这样s1 s2都是有效的 
如果一个类型实现了copy trait 旧的变量在赋值之后仍然可以使用




## 生命周期

'a 生命周期声明的含义

>Hi, I found your answer not simple enough for Rust newbies like me. After looking around for quite some time, I'd like to offer a simpler explanation to the question: "Why is 'a there?" ---- Because the struct definition ties it to a referenced object (in this case, every struct Person instance is referencing a &str) you want to specificly declare an arbitary lifetime and tie these two things together: You want a struct Person instance to only live as long as its referenced object (hence Person<'a> and name: &'a str) so dangling references after each other's death is avoided. 

```rust
#[derive(Debug)]
struct Person<'a> {
    name: &'a str,
    age: u8
}

fn main() {
    let name = "Peter";
    let age = 27;
    let peter = Person { name, age };

    // Pretty print
    println!("{:#?}", peter);
}
 
```

- https://stackoverflow.com/questions/47640550/what-is-a-in-rust-language





## cargo
cargo下载依赖的命令

在 Rust 的项目中，使用 cargo 下载依赖的命令是：

cargo build

这个命令会读取 Cargo.toml 文件中的依赖关系，并从 crates.io 下载依赖的库，最后编译整个项目。



下载的依赖包的存放位置

下载的依赖包存放在 $HOME/.cargo/registry 目录下面，这个目录是用户级别的，因此可以被所有项目共享。

你也可以在某个项目目录下面看到本地编译生成的文件，它们存放在 target 目录下面。

需要注意的是，如果你使用了不同的 Rust 版本，那么你可能会在 $HOME/.cargo 目录下面看到不同的文件夹，例如：$HOME/.cargo/registry/src 和 $HOME/.cargo/registry/{hash}/src。这是因为 cargo 会根据 Rust 的版本为不同的项目生成不同的依赖关系。


- https://course.rs/about-book.html