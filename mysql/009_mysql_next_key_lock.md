id	age	name
1	10	Lee
3	24	Soraka
5	32	Zed
7	45	Talon

该表中 age 列潜在的临键锁有：
(-∞, 10],		
(10, 24],		
(24, 32],		
(32, 45],		
(45, +∞],

在事务 A 中执行如下命令：
-- 根据非唯一索引列 UPDATE 某条记录
UPDATE table SET name = Vladimir WHERE age = 24;
-- 或根据非唯一索引列 锁住某条记录
SELECT * FROM table WHERE age = 24 FOR UPDATE;
事务 A 在对 age 为 24 的列进行 UPDATE 操作的同时，也获取了 (24, 32] 这个区间内的临键锁。
这样就可以解决幻读问题。。



不管执行了上述 SQL 中的哪一句，之后如果在事务 B 中执行以下命令，会被阻塞：
INSERT INTO table VALUES(id=100, age=26, name='Ezreal');
INSERT INTO table VALUES(100, 30, 'Ezreal');

