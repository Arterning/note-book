/**
 * 除了直接用{ ... }创建一个对象外，JavaScript还可以用一种构造函数的方法来创建对象。它的用法是，先定义一个构造函数：
 * 
 * 你会问，咦，这不是一个普通函数吗？
这确实是一个普通函数，但是在JavaScript中，可以用关键字new来调用这个函数，并返回一个对象

注意，如果不写new，这就是一个普通函数，它返回undefined。
但是，如果写了new，它就变成了一个构造函数，它绑定的this指向新创建的对象，并默认返回this，也就是说，不需要在最后写return this;。

原型链是：
xiaoming ----> Student.prototype ----> Object.prototype ----> null

 */
function Student(name) {
  this.name = name;
  this.hello = function () {
    alert("Hello, " + this.name + "!");
  };
}

/**
 * 用new Student()创建的对象还从Student.prototype原型上获得了一个constructor属性，它指向函数Student本身：
 */
xiaoming.constructor === Student.prototype.constructor; // true

/****************************************************************************************/

/**
 * xiaoming和xiaohong各自的hello是一个函数，但它们是两个不同的函数，虽然函数名称和代码都是相同的！
如果我们通过new Student()创建了很多对象，这些对象的hello函数实际上只需要共享同一个函数就可以了，这样可以节省很多内存。

 */
xiaoming.name; // '小明'
xiaohong.name; // '小红'
xiaoming.hello; // function: Student.hello()
xiaohong.hello; // function: Student.hello()
xiaoming.hello === xiaohong.hello; // false
/**
 *我们只要把hello函数移动到xiaoming、xiaohong这些对象共同的原型上就可以了，也就是Student.prototype
 感觉prototype就是class
 */
function Student(name) {
  this.name = name;
}

Student.prototype.hello = function () {
  alert("Hello, " + this.name + "!");
};
