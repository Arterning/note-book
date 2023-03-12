/**
 * 全局匹配可以多次执行exec()方法来搜索一个匹配的字符串。当我们指定g标志后，每次运行exec()，正则表达式本身会更新lastIndex属性，表示上次匹配到的最后索引：
用+表示至少一个字符，用?表示0个或1个字符

正则表达式还可以指定i标志，表示忽略大小写，m标志，表示执行多行匹配。都是放在最后的

 */
var s = "JavaScript, VBScript, JScript and ECMAScript";
var re = /[a-zA-Z]+Script/g;

// 使用全局匹配:
re.exec(s); // ['JavaScript']
re.lastIndex; // 10

re.exec(s); // ['VBScript']
re.lastIndex; // 20

re.exec(s); // ['JScript']
re.lastIndex; // 29

re.exec(s); // ['ECMAScript']
re.lastIndex; // 44

re.exec(s); // null，直到结束仍没有匹配到
