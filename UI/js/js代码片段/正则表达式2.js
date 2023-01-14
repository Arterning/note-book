/**
 * JavaScript有两种方式创建一个正则表达式：
第一种方式是直接通过/正则表达式/写出来，
第二种方式是通过new RegExp('正则表达式')创建一个RegExp对象。
注意，如果使用第二种写法，因为字符串的转义问题，字符串的两个\\实际上是一个\。

 */

function base() {
  var re1 = /ABC\-001/;
  var re2 = new RegExp("ABC\\-001");

  re1; // /ABC\-001/
  re2; // /ABC\-001/

  //先看看如何判断正则表达式是否匹配：

  var re = /^\d{3}\-\d{3,8}$/;
  re.test("010-12345"); // true
  re.test("010-1234x"); // false
  re.test("010 12345"); // false
}
