闭包就是在函数外部访问函数局部变量的方式

f2可以读取f1中的局部变量，那么只要把f2作为返回值，我们就可以在f1外部读取它的内部变量了

```css
function f1(){

　　　　var n=999;

　　　　function f2(){
　　　　　　alert(n);
　　　　}

　　　　return f2;

　　}

　　var result=f1();

　　result(); // 999
```

闭包可以用在许多地方。它的最大用处有两个，一个是前面提到的可以读取函数内部的变量，另一个就是让这些变量的值始终保持在内存中。

```css
   function f1() {

　　　　var n = 999;

　　　　add = function(){n+=1}

　　　　function f2(){
　　　　　　alert(n);
　　　　}

　　　　return f2;

　　}

　　var result = f1();

　　result(); // 999

　　add();

　　result(); // 1000
```

在这段代码中，result实际上就是闭包f2函数。它一共运行了两次，第一次的值是999，第二次的值是1000。这证明了，函数f1中的局部变量n一直保存在内存中，并没有在f1调用后被自动清除。

原因就在于f1是f2的父函数，而f2被赋给了一个全局变量result，这导致f2始终在内存中，而f2的存在依赖于f1，因此f1也始终在内存中，不会在调用结束后，被垃圾回收机制（garbage collection）回收。

如果你能理解下面两段代码的运行结果，应该就算理解闭包的运行机制了。

代码片段一。

```jsx
var name = "The Window";

　　var object = {
　　　　name : "My Object",

　　　　getNameFunc : function(){
　　　　　　return function(){
　　　　　　　　return this.name;
　　　　　　};

　　　　}

　　};

//object.getNameFunc()返回闭包函数
//接着再直接调用闭包函数，所以闭包函数的this就是window
alert(object.getNameFunc()());//The Window
```

代码片段二。

```jsx
var name = "The Window";

　　var object = {
　　　　name : "My Object",

　　　　getNameFunc : function(){
　　　　　　var that = this;
　　　　　　return function(){
　　　　　　　　return that.name;
　　　　　　};

　　　　}

　　};

　　alert(object.getNameFunc()());//My Object
```

```jsx
var name = "The Window";

　　var object = {
　　　　name : "My Object",

　　　　getNameFunc : function(){
　　　　　　return () => {
　　　　　　　　return this.name;//箭头函数继承外层this
　　　　　　};

　　　　}

　　};

　　alert(object.getNameFunc()());//My Object
```