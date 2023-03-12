/**
 * 传统的JavaScript程序使用函数和基于原型的继承来创建可重用的组件，但对于熟悉使用面向对象方式的程序员来讲就有些棘手，因为他们用的是基于类的继承并且对象是由类构建出来的。 从ECMAScript 2015，也就是ECMAScript 6开始，JavaScript程序员将能够使用基于类的面向对象的方式。 使用TypeScript，我们允许开发者现在就使用这些特性，并且编译后的JavaScript可以在所有主流浏览器和平台上运行，而不需要等到下个JavaScript版本。
 *
 * 对于我这种从java转过来的来说 使用typescript写前端是非常合适的，而且现在typescript有类型检查 目前也逐渐成为前端事实的标准了 学习他性价比还是比较高的
 * 学东西要讲究性价比 有的东西可以根据爱好看一眼 但是如果一直死磕的话 其实对你日后涨薪可能没什么帮助 还蹉跎了岁月 没有必要 先把工资涨上去再说 想让自己过上想要的生活再说
 * 怎么做？很明确了 学习TS 学习前端和nodeJS
 */
class Greeter {
  greeting: string;
  constructor(message: string) {
    this.greeting = message;
  }
  greet() {
    return "Hello, " + this.greeting;
  }
}

let greeter = new Greeter("world");
