/*********************************************************************
rust核心特性是所有权
rust没有垃圾搜集器
在编译的阶段就做好了内存管理的工作

Stack 栈内存
存储在Stack上的数据必须是已知的固定的大小
分配非常快
当函数被调用的时候
函数的本地变量会被压入Stack 函数调用结束后弹出这些变量

Heap 堆内存
存储在Heap上的数据的内存空间是未知的
**/



/*********************************************************************
所有权解决的问题
1. 跟踪代码的哪些部分正在使用heap的哪些数据
2. 最小化heap的重复数据量
3. 清理heap上的未使用的数据

String类型是存储在Heap上 因为他是可变的
字符串字面值是硬编码写入可执行文件的

当拥有某个值的变量走出他的作用域以后，就自动释放内存
**/





/*********************************************************************
移动
let x =5;
let y = x;
x y都是有效的 因为x是整数类型的 对于这种类型的值的复制操作是非常快速的
整数是已知大小的值 所有两个5都被压入stack中

所有权转移
let s1=String::From("hello")
let s2=s1
//s1被废弃 无法再被使用
//rust不会自动创建对象的深拷贝 如果需要对heap上的对象进行deep clone 需要使用clone()

克隆
let s1=String::From("hello")
let s2=s1.clone()
这样s1 s2都是有效的 
如果一个类型实现了copy trait 旧的变量在赋值之后仍然可以使用
**/




/*
// 01. take ownership
fn main() {
    let s = String::from("hello");
    takes_ownership(s);s传过去之后就不再有效了

    let x = 5;
    make_copy(x);传递的是x的副本
}

fn takes_ownership(some_string: String) {
    println!("{}", some_string);
}

fn make_copy(some_integer: i32) {
    println!("{}", some_integer);
}
*/



/*// 02. give ownership
fn main() {
    let s1 = give_onwership();

    let s2 = String::from("hello");
    let s3 = takes_and_give_back(s2);

    println!("s1: {}; s3:{}", s1, s3);
}

//字符串的所有权从some_string转移到了main中的s1
fn give_onwership() -> String {
    let some_string = String::from("hello");

    some_string
}

//拿来又送回去
fn takes_and_give_back(a_string: String) -> String {
    a_string
}
*/


/*// 03. return tuple 借用的第一种方式 为了归还所有权又返回s2 这种做法比较麻烦
fn main() {
    let s1 = String::from("hello");

    let (s2, len) = calculate_length(s1);

    println!("The length of '{}' is {}.", s2, len);
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len();

    (s, length)
}
*/



/*// 04. borrowing 让函数使用某个值 但是不获得所有权，使用引用类型
引用就是允许你使用某些值 但是不获得所有权 这种方式就是借用
我们是否可以修改借用的东西呢？答案是不行的 不可以修改借用的东西
引用默认情况下是不可变的 需要加上mut关键字
fn main() {
    let s1 = String::from("hello");

    let len = calculate_length(&s1);

    println!("The length of '{}' is {}.", s1, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}
*/

/**
//可变引用
fn main() {
    let mut s = String::from("hello");

    change(&mut s);

    println!("s: {}", s);
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
**/


/*05. 
下面的代码是有问题的 只能有一个可变的引用
这样做的好处是在编译的阶段就防止数据竞争
1. 两个或者多个指针访问同一块内存
2. 没有使用同步机制
3. 至少有一个指针在写入数据
就会发生数据竞争 在编译阶段就会报错
rust做了一个根本的解决就是只能有一个可变引用
fn main() {
    let mut s = String::from('hello');
    let s1 = &mut s;
    let s2 = &mut s;
}
*/



/*06.
我们可以在不同的作用域隔开这些可变引用就可以了
fn main() {
    let mut s = String::from('hello');
    {
        let s1=&mut s;
    }
    let s2 = &mut s;

}

*/