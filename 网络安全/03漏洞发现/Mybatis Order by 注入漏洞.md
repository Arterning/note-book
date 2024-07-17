#{}相当于jdbc中的preparedstatement 带引号

${}是输出变量的值 不带引号



1)order by后面如果采用预编译的形式动态输入参数,那么实际插入的参数是一个字符串,例子中是:order by 'domain_id'

2)输出结果并没有排序,从sql语句中的形式我们也可以推测出此sql语句根本也不合法(正常应该是 order by domain_id)



这里先说一下只能${}的情况,从我们前面的例子中也能看出,order by是肯定只能用${}了,用#{}会多个' '导致sql语句失效.此外还有一个like 语句后也需要用${},简单想一下

就能明白.由于${}仅仅是简单的取值,所以以前sql注入的方法适用此处,如果我们order by语句后用了${},那么不做任何处理的时候是存在sql注入危险的.你说怎么防止,那我只

能悲惨的告诉你,你得手动处理过滤一下输入的内容,如判断一下输入的参数的长度是否正常(注入语句一般很长),更精确的过滤则可以查询一下输入的参数是否在预期的参数集合中..



你说怎么防止,那我只

能悲惨的告诉你,你得手动处理过滤一下输入的内容,如判断一下输入的参数的长度是否正常(注入语句一般很长),更精确的过滤则可以查询一下输入的参数是否在预期的参数集合中..



比如order by 只能是你指定的那几个列名。。


```java
private static String badStr = "'|and |exec |execute |insert |select |delete |update |drop |* |% |chr |mid |master |truncate |declare |sitename |net user |xp_cmdshell |;|or |-|+|like'|create |table |from |grant |use |group_concat |column_name |information_schema.columns|table_schema|union |where |char |--|like |//|/|#";
    private static final String[] keywords = new String[]{"'", "and", "exec", "execute", "insert", "select", "delete", "update", "drop", "*", "%", "chr", "mid", "master", "truncate", "declare", "sitename", "net user", "xp_cmdshell", ";", "or", "-", "+", "like'", "create", "table", "from", "grant", "use", "group_concat", "column_name", "information_schema.columns", "table_schema", "union", "where", "char", "--", "like", "//", "/", "#"};
    private static final Pattern CREATE_DATE_PATTERN = Pattern.compile("([a-zA-Z][a-zA-Z0-9_]{0}\\.)?create_date");

    public SecurityUtils() {
    }

    public static boolean isSqlSafe(String sql) {
        if (StringUtils.isNotBlank(sql)) {
            String replace = sql.toLowerCase().trim();
            Matcher matcher = CREATE_DATE_PATTERN.matcher(replace);
            if (matcher.find()) {
                return true;
            }

            String[] var3 = keywords;
            int var4 = var3.length;

            for(int var5 = 0; var5 < var4; ++var5) {
                String keyword = var3[var5];
                if (replace.indexOf(keyword) >= 0) {
                    return false;
                }
            }
        }

        return true;
    }
```



直接在一个项目中搜索·`order by ${orderBy}`可以搜到很多这种


```xml
        <if test="orderBy != null and orderBy != ''">
            order by ${orderBy}
            <if test="order != null and order != ''">
                ${order}
            </if>
        </if>
        <if test="orderBy == null or orderBy == ''">
            ORDER BY a.create_date desc
        </if>
```