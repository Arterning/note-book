## 普通函数的this
普通函数的this 和java中的this是一样的，就是函数的调用者

在箭头函数出现之前，每一个新函数根据它是被如何调用的来定义这个函数的 this 值：

- 如果该函数是一个构造函数，this 指针指向一个新的对象
- 在严格模式下的函数调用下，this 指向undefined
- 如果该函数是一个对象的方法，则它的 this 指针指向这个对象

This被证明是令人厌烦的面向对象风格的编程。

```javascript
function Person() {
  // Person() 构造函数定义 `this`作为它自己的实例。
  this.age = 0;

  setInterval(function growUp() {
    // 在非严格模式，growUp() 函数定义 `this`作为全局对象，
    // 与在 Person() 构造函数中定义的 `this`并不相同。
    this.age++;
  }, 1000);
}

var p = new Person();
```
在上面这个例子中，this.age 并不指向Person中的age
因为groupUp回调函数并不是通过Person对象调用的


## 箭头函数的this

**箭头函数没有自己的this指向**，它的this指向上一级作用域的this

下面这个例子setInterval中的箭头函数的上一层是Person，因为他上一层有一个函数Person,所以这个外层，指的是函数。

箭头函数会继承外层函数的this值

```javascript
function Person(){
	this.age = 0;

	setInterval(() => {
		this.age++; // |this| 正确地指向 p 实例
	}, 1000);
}

var p = new Person();
```

如果箭头函数外面没有包裹函数，那么它就是指向this

```javascript
const obj = {
    a: () => {
        console.log(this);
    }
}
obj.a();
```