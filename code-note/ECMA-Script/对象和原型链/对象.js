/**
 * 当我们用obj.xxx访问一个对象的属性时，JavaScript引擎先在当前对象上查找该属性，如果没有找到，就到其原型对象上找，如果还没有找到，就一直上溯到Object.prototype对象，最后，如果还没有找到，就只能返回undefined。
 */

/**
 * 创建一个Array对象：
 * 其原型链是：
    arr ----> Array.prototype ----> Object.prototype ----> null

    Array.prototype定义了indexOf()、shift()等方法，因此你可以在所有的Array对象上直接调用这些方法。
 */
var arr = [1, 2, 3];

/**
 * 当我们创建一个函数时：
 * 函数也是一个对象，它的原型链是：
foo ----> Function.prototype ----> Object.prototype ----> null
由于Function.prototype定义了apply()等方法，因此，所有函数都可以调用apply()方法。

很容易想到，如果原型链很长，那么访问一个对象的属性就会因为花更多的时间查找而变得更慢，因此要注意不要把原型链搞得太长。
 */
function foo() {
  return 0;
}

/**
 * __proto__是什么
 */
