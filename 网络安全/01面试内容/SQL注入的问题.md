1. SQL
SQL注入原理
把用户的输入作为sql代码执行了

SQL注入的种类?
注入点类型分类 数字型注入 字符型 搜索型注入
按照数据提交方式 get注入 post注入 http头部注入
按照执行效果 盲注 联合查询 堆叠注入 宽字节注入 二次注入

SQL注入的绕过有了解过吗?有绕过那些wAF?绕wAF的方法有哪些?


介绍一下你的SQL注入流程
直接输入一个闭合符号判断有没有注入漏洞 如果报错了 肯定有 直接放到sqlmap里面扫描


SQL注入一开始用来判断数据库类型的语句是什么?
1. 构造一个错误的sql 如果后端没有catch异常的话 可以观察错误信息
2. 通过特有数据表判断 比如mysql5以后特有的information_schema
3. 基于特定函数的判断 len和length 在mssql和mysql以及db2内，返回长度值是调用len()函数；在oracle和INFORMIX则是通过length()来返回长度值。

讲一下sql写入webshell的条件，如果你在的目录没有执行权限怎么办
secure_file_priv 这个变量被用于限制导入和导出的数据目录
secure_file_priv 有些设置选项:
如果为空，不做目录限制，即任何目录均可以。
如果指定了目录，MySQL 会限制只能从该目录导入、或导出到该目录。目录必须已存在，MySQL 不会自动创建该目录。
如果设置为 NULL，MySQL 服务器禁止导入与导出功能。

MySQL中root 用户拥有所有权限，但写入Webshell并不需要一定是root用户权限，比如数据库用户只要拥有FILE权限就可以执行 select into outfile操作。

在mysql的配置文件 my.ini 中 也可以配置secure_file_priv="c:/wamp64/tmp" 

当secure_file_priv文件导出路径与web目录路径重叠，写入Webshell才可以被访问到。
如果所在目录没有执行权限 那么只能修改secure_file_priv的配置了 修改成有执行权限的目录




讲讲sql注入的条件
目标有sql漏洞


SQL注入里面有一个into outfile函数有什么作用?
导出表里面的数据到文件


xp_cmdshell 有了解过吗?能用来做什么
MSSQL的一个存储过程漏洞
https://www.jianshu.com/p/9d5e8236647e


你有了解过那些数据库的sql注入?你是怎么快速判断它属于那个数据库的呢?
判断数据库类型 可以看他开放的端口
比如mysql默认端口是3306 oracle是15201
如果运维没有修改这个默认端口 基本可以猜测出来


时间延时注入的原理?sleep ()函数不能用你一般会怎么处理
可以用布尔值盲注
原理就是在if里面构造猜测语句 如果满足条件就sleep()
kobe' and if((substr(database(),1,1))='a',sleep(5),null)


HTTP头注入了解过吗?注入点一般存在于那些地方
比如后端程序需要读取http头信息作为参数，如果这个参数传递给了数据库的话就是头注入
总之关键是后台的参数是怎么传递的。这个参数可以广泛的分布在web应用程序的任何地方
比如UserAgent


SQL注入的防护方法有哪些呢?预编译的原理是什么呢?
WAF 预编译
数据库接受到sql语句之后，需要词法和语义解析，优化sql语句，制定执行计划。这需要花费一些时间。
预编译语句就是将这类语句中的值用占位符替代，可以视为将sql语句模板化或者说参数化。一次编译、多次运行，省去了解析优化等过程。


orderby注入了解吗
经常利用order by子句进行快速猜解表中的列数
order by后面的参数无法预编译
