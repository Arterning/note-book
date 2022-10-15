方法调用和返回

方法可以分为两类：
静态方法（或者类方 法）
实例方法。

静态方法通过类来调用，实例方法则通过对象引 用来调用

Java虚拟机规范一共提供了4条方法调用指令
invokestatic指令用来调用静态方法
invokespecial指令用来调 用无须动态绑定的实例方法
如果是针对 接口类型的引用调用方法，就使用invokeinterface指令
invokevirtual