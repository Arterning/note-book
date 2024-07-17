
原因： 问题的出现是因为MySQL对于BLOB、TEXT、GEOMETRY和JSON字段是不允许有默认值的，在mysql5.7以后就有了严格模式`sql_model`规定了此限制。

1. 查询sql_mode,

```sql
show variables like '%sql_mode%';
```

2. `STRICT_TRANS_TABLES`即是导致上面报错产生的原因，去掉它即可。




原因在于：

1、  MYSQL5.x是不允许BLOB/TEXT类型的字段拥有默认值的。

2、  由于MYSQL是在‘strict mode’严格模式下工作的，如果改为非严格模式，上面的语句就可以执行成功

3、  MYSQL5.x在windows下是默认以‘strict mode’工作的，当执行上面的语句时，会给你一个错误或者警告信息

解决方法：

1、  找到mysql安装根目录下的my.ini文件

2、  找到这样一行：  

```
sql-mode="STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"
```

3、  在其前面加‘#’将其注释掉：

```
#sql-mode="STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"
```

4、  重启mysql服务

5、  重新执行你的mysql语句