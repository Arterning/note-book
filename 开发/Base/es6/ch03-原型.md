显式原型：prototype 属于方法 Funtion.prototype
隐式原型：__ proto__ 属于对象 obj.__proto__

显式原型和隐式原型是什么？
在js中万物皆对象，方法(Function)是对象，方法的原型（Function.prototype）是对象，对象具有属性（__ proto__）称为隐式原型，对象的隐式原型指向构造该对象的构造函数的显式原型 （prototype）。

方法(Function)是一个特殊的对象，除了和其他对象一样具有__proto__属性以外，它还有一个自己特有的原型属性(prototype)，这个属性是一个指针，指向原型对象。原型对象也有一个属性叫constructor，这个属性包含一个指针，指向原构造函数。

注意：通过Function.prototype.bind方法构造出来的函数没有prototype属性。
注意：Object.prototype.这个对象的是个例外，它的__proto__值为null。

二者的关系
隐式原型指向创建这个对象的构造函数的prototype

二者的区别
prototype:
专属于函数的一个属性，类型为对象，叫原型对象；
作用：为了给将来自身所在的构造函数被new出来的实例做父级使用的
__ proto__：
专属于对象数据的一个属性，类型为对象，叫隐式原型，
作用：找父级
特性：

a.当某个对象自身不具有某个属性或方法时，会找父级
b.当这个对象是被new出来的实例时，这个对象的父级（__ proto__）就是当前被new出来的构造函数的prototype