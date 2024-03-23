

## MyBatis中 #{} 和 ${} 区别


## #{} 方式

#{}: 解析为SQL时，会将形参变量的值取出，并自动给其添加引号。 例如：当实参username="Amy"时，传入下Mapper映射文件后

```xml
    <select id="findByName" parameterType="String" resultMap="studentResultMap">
        SELECT * FROM user WHERE username=#{value}
    </select>
```


这样感觉#{} 方式也不是绝对安全，

如果攻击者构造这样的输入呢？

```sql
ss";select * from users;--
```

那么最终的结果岂不是就会变成

```sql

select .. from table where username = "ss";select * users; --"
```


岂不是还是会造成SQL注入？



## ${} 方式

${}: 解析为SQL时，将形参变量的值直接取出，直接拼接显示在SQL中




## 总结

#{}方式则是先用占位符代替参数将SQL语句先进行预编译，然后再将参数中的内容替换进来。由于SQL语句已经被预编译过，其SQL意图将无法通过非法的参数内容实现更改，其参数中的内容，无法变为SQL命令的一部分。故，**#{}可以防止SQL注入而${}却不行**

所以我感觉她的说法有问题，#{} 并不是自动给其添加引号这么简单，她的本质是SQL的预编译。

换句话说，假如攻击者的输入是

```
ss";select * from users;--
```

那么也会变成

```sql
SELECT * FROM todo WHERE name = 'ss'';select * from users;--'
```

也就是说攻击者构造的"" 会被转义成''


如果攻击者的输入是

```sql
ss';select * from users;--
```

那么也单引号也会被转移成''

```sql
SELECT * FROM todo WHERE name = 'ss'';select * from users;--'
```


所以预编译的sql在替换的时候，对于特殊字符都会进行转义，保证用户构造的任何输入一定是在一个单引号之中！



## 只能使用${}的场景

由于#{}会给参数内容自动加上引号，会在有些需要表示字段名、表名的场景下，SQL将无法正常执行。

期望查询结果按sex字段升序排列，参数String orderCol = "sex",mapper映射文件使用#{}：

```xml
<select id="findAddByName3" parameterType="String" resultMap="studentResultMap">
        SELECT * FROM USER WHERE username LIKE '%Am%' ORDER BY #{value} ASC
    </select>
```


