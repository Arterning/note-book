# PostgreSQL 支持以下数据类型：

基本类型：
Smallint: 2 字节整数
Integer: 4 字节整数
Bigint: 8 字节整数
Real: 单精度浮点数
Double Precision: 双精度浮点数
Serial: 自增的 4 字节整数
BigSerial: 自增的 8 字节整数
Numeric: 可变精度数字
Decimal: 可变精度数字，等价于 Numeric

字符串类型：
Character(n) / Char(n): 固定长度字符串，最大长度为 n
Varchar(n): 可变长度字符串，最大长度为 n
Text: 变长字符串，不限长度

二进制类型：
Bytea: 二进制字符串

日期和时间类型：
Date: 日期
Time: 时间
Timestamp: 日期和时间
Timestamp with time zone: 带时区的日期和时间
Interval: 时间间隔
布尔类型：
Boolean: 布尔值
UUID 类型：
UUID: Universally Unique Identifier，全局唯一标识符

地理空间类型：
Point: 点
Line: 线
Lseg: 线段
Box: 矩形
Path: 路径
Polygon: 多边形
Circle: 圆
其他类型：
Bit: 位串
Bit Varying: 变长位串
MACADDR: MAC 地址
INET: IP 地址
CIDR: IP 地址块

# postgresql 定义主键一般用什么类型的字段？

在 PostgreSQL 中，一般使用整数类型的字段来定义主键。常用的整数类型有 "integer" 和 "serial"，其中 "serial" 是一个自动增长的整数类型，比较方便使用。
这个概念类似于 oracle 的 sequence 自增序列

"serial" 类型字段使用 4 字节整数存储，范围为 -2147483648 到 2147483647。

如果你的应用需要存储更大的数据范围，那么可以使用 "bigserial" 类型，它使用 8 字节整数存储，范围为 -9223372036854775808 到 9223372036854775807。
