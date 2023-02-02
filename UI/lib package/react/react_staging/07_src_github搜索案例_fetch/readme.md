fetch不依赖xhr 是内置的
不是第三方库 不需要安装 而且也是Promise风格

那么缺点是什么呢？


then()的返回是依然是Promise对象 这也是可以链式调用的原因
Promise.all(result1, result2)

这一节非常重要
对于Promise的理解
async awiat关键字的使用非常重要