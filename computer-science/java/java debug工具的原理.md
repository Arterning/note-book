class文件中的LineNubmberTable 描述了java源代码和字节码行号之间的对应关系





JVM如果要能够调试

JVM规范定义JVM需要实现这些功能

- 获取一个线程的状态， 挂起一个线程，让线程恢复执行， 设置一个线程，单步执行
- 获取线程的当前栈帧，调用栈帧，栈帧对应的方法名
- 获取变量的值， 设置变量的值
- 设置断点，清除断点
- 查看类的信息，方法，字段 等等”



 JVM提供C语言的接口。嗯，这个接口就叫做JVM Tool Interface，简称JVM TI。” 

 协议，供调试器和JVM 通信 ！这个协议的名称就叫 （JDWP）Java Debug Wire Protocol  



 提供一个Java版本的接口吧，让这个接口把JDWP还有什么JVM TI都给封装起来，主要供我们的Java IDE来使用，来集成。叫做JVM TI



 “嗯，层次划分得不错，程序员可以直接使用JVM 提供的接口，也可以用JDWP， 还可以用JDI..... ” 





![微信截图_20200101165524](..\images\微信截图_20200101165524.png)





```java
ClassPrepareEvent event = ....略....
ClassType classType = (ClassType) event.referenceType();
// 获取表示第10行的Location对象
Location location = classType.locationsOfLine(10).get(0);
// 通过Location对象创建一个断点
BreakpointRequest bpReq = vm.eventRequestManager().createBreakpointRequest(location);
bpReq.enable();  
```



```java
// 到达了一个断点
BreakpointEvent event = .....略......
//获取当前的线程
ThreadReference threadReference = event.thread();
//获取当前的栈帧
StackFrame stackFrame = threadReference.frame(0);
//从栈帧中得到本地变量 i
LocalVariable localVariable = stackFrame.visibleVariableByName("i");
Value value = stackFrame.getValue(localVariable);
int i = ((IntegerValue) value).value();
System.out.println("The local variable " + "i" + " is " + i);  
```

