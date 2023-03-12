/**
 * 
 * JSON是JavaScript Object Notation的缩写，它是一种数据交换格式。

在JSON出现之前，大家一直用XML来传递数据。因为XML是一种纯文本格式，所以它适合在网络上交换数据。XML本身不算复杂，但是，加上DTD、XSD、XPath、XSLT等一大堆复杂的规范以后，任何正常的软件开发人员碰到XML都会感觉头大了，最后大家发现，即使你努力钻研几个月，也未必搞得清楚XML的规范。

终于，在2002年的一天，道格拉斯·克罗克福特（Douglas Crockford）同学为了拯救深陷水深火热同时又被某几个巨型软件企业长期愚弄的软件工程师，发明了JSON这种超轻量级的数据交换格式。


被某几个巨型软件企业长期愚弄 说的真的对 我也觉得XML真的不好用
我还是喜欢JSON
所以以后转前端吧 配置文件都是JSON 看起来就美观
不像JAVA 以前的各种spring struts hibernate在没有注解的时候 配置文件都是XML 看着就不怎么样


在JSON中，一共就这么几种数据类型：

number：和JavaScript的number完全一致；
boolean：就是JavaScript的true或false；
string：就是JavaScript的string；
null：就是JavaScript的null；
array：就是JavaScript的Array表示方式——[]；
object：就是JavaScript的{ ... }表示方式。
并且，JSON还定死了字符集必须是UTF-8，表示多语言就没有问题了。为了统一解析，JSON的字符串规定必须用双引号""，Object的键也必须用双引号""。

由于JSON非常简单，很快就风靡Web世界，并且成为ECMA标准。几乎所有编程语言都有解析JSON的库，而在JavaScript中，我们可以直接使用JSON，因为JavaScript内置了JSON的解析。

把任何JavaScript对象变成JSON，就是把这个对象序列化成一个JSON格式的字符串，这样才能够通过网络传递给其他计算机。
如果我们收到一个JSON格式的字符串，只需要把它反序列化成一个JavaScript对象，就可以在JavaScript中直接使用这个对象了。
 */

var xiaoming = {
  name: "小明",
  age: 14,
  gender: true,
  height: 1.65,
  grade: null,
  "middle-school": '"W3C" Middle School',
  skills: ["JavaScript", "Java", "Python", "Lisp"],
};
var s = JSON.stringify(xiaoming);

//第二个参数用于控制如何筛选对象的键值，如果我们只想输出指定的属性，可以传入Array：
JSON.stringify(xiaoming, ["name", "skills"], "  ");

//还可以传入一个函数，这样对象的每个键值对都会被函数先处理：
// 下面的代码把所有属性值都变成大写：
function convert(key, value) {
  if (typeof value === "string") {
    return value.toUpperCase();
  }
  return value;
}

JSON.stringify(xiaoming, convert, "  ");

/**
 * 如果我们还想要精确控制如何序列化小明，可以给xiaoming定义一个toJSON()的方法，直接返回JSON应该序列化的数据：
 */
var xiaoming = {
  name: "小明",
  age: 14,
  gender: true,
  height: 1.65,
  grade: null,
  "middle-school": '"W3C" Middle School',
  skills: ["JavaScript", "Java", "Python", "Lisp"],
  toJSON: function () {
    return {
      // 只输出name和age，并且改变了key：
      Name: this.name,
      Age: this.age,
    };
  },
};
