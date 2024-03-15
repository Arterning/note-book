

es6类写法

```js
    class Person {
        constructor(name) {
            this.name = name;
        }
        printName() {
            console.log('This is printName');
        }
        commonMethods(){
            console.log('我是共享方法');
        }
    }

    class Student extends Person {
        constructor(name, score) {
            super(name);
            this.score = score;
        }
        printScore() {
            console.log('This is printScore');
        }
    }

    let stu = new Student('小红');
    let person = new Person('小紫');
    console.log(stu.commonMethods===person.commonMethods);//true
```


原生写法

```java
    function Person (name){
        this.name = name;
        this.printName=function() {
            console.log('This is printName');
        };
    }
    Person.prototype.commonMethods=function(){
        console.log('我是共享方法');
    };

    function Student(name, score) {
        Person.call(this,name);
        this.score = score;
        this.printScore=function() {
            console.log('This is printScore');
        }
    }
    Student.prototype = new Person();
    let person = new Person('小紫',80);
    let stu = new Student('小红',100);
    console.log(stu.printName===person.printName);//false
    console.log(stu.commonMethods===person.commonMethods);//true
```

可以看到es6写法还是比原生写法简洁许多的



**构造函数特点：**

1.构造函数有原型对象prototype。

2.构造函数原型对象prototype里面有constructor，指向构造函数本身。

3.构造函数可以通过原型对象添加方法。

4.构造函数创建的实例对象有**proto**原型，指向构造函数的原型对象。




**类：**

1.class本质还是function

2.类的所有方法都定义在类的prototype属性上

3.类创建的实例，里面也有**proto**指向类的prototype原型对象

4.新的class写法，只是让对象原型的写法更加清晰，更像面向对象编程的语法而已。

5.ES6的类其实就是语法糖。


什么是语法糖？加糖后的代码功能与加糖前保持一致，糖在不改变其所在位置的语法结构的前提下，实现了运行时的等价。

语法糖没有改变语言功能，但增加了程序员的可读性。








原文：
https://cloud.tencent.com/developer/article/2226846
