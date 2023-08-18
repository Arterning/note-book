Lambda表达式
把一块代码赋给一个变量


aBlockOfCode = (s）-> System.out.println(s);
s的参数类型是String，编译器可以自动判断
方法没有返回值。

函数式接口
@FunctionalInterface
interface MyLambdaInterface {
  void doSomeShit(String s);
}
@FunctionalInterface 可以防止其他人往这个接口里面添加新的函数。
那么，一个完整的lambda表达式声明就是
MyLambdaInterface aBlockOfCode = (s) -> System.out.println(s);

以前，没有lambda表达式，需要new一个实现了这个接口的类，然后覆盖这个方法。我们传入这个类。
Thread thread = new Thread(new Runnable() {
    @Override
    public void run() {
        System.out.println("hi");
    }
});

现在，有了lambda表达式，我们只需要把run()方法里面的这一段代码传过去就可以了。
Thread thread = new Thread(() -> {
            System.out.println("hi");
        });
thread.start();
