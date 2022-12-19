Java 8 API添加了一个新的抽象称为流Stream，可以让你以一种声明的方式处理数据。
Stream 使用一种类似用 SQL 语句从数据库查询数据的直观方式来提供一种对 Java 集合运算和表达的高阶抽象。
Stream API可以极大提高Java程序员的生产力，让程序员写出高效率、干净、简洁的代码。


1. 生成流
list.stream() 串行流
list.parallerStream() 并行流

2. forEarch
Stream 提供了新的方法 'forEach' 来迭代流中的每个数据。以下代码片段使用 forEach 输出了10个随机数：
Random random = new Random();
random.ints().limit(10).forEach(System.out::println);


3. 聚合操作 类似SQL语句一样的操作， 比如filter, map, reduce, find, match, sorted等。
List<Integer> squaresList = numbers.stream().map( i -> i*i).distinct().collect(Collectors.toList());
//过滤掉空的字符串
int count = strings.stream().filter(string -> string.isEmpty()).count();
limit 方法用于获取指定数量的流。 以下代码片段使用 limit 方法打印出 10 条数据：
random.ints().limit(10).forEach(System.out::println);

Collectors
Collectors 类实现了很多归约操作，例如将流转换成集合和聚合元素。Collectors 可用于返回列表或字符串：
List<String> filtered = strings.stream().filter(string -> !string.isEmpty()).collect(Collectors.toList());
