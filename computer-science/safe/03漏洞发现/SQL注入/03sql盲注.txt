后台使用了错误信息屏蔽方法屏蔽了数据库的报错信息
也就是java catch了异常
这种情况下的注入 就是盲注
分为based boolean 和 based time


kobe' and 1=1#
kobe' and 1=2#


猜字符
select ascii(substr(database(),1,1)) > 100;
结果为1或者0
我们可以用这个结果替代and 1=1 / and 1=0
一个字符一个字符的比较 就是比较慢比较麻烦
还是用sqlmap工具


猜长度
select length(database()) > 10;