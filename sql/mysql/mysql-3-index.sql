CREATE INDEX indexName ON table_name (column_name)
DROP INDEX [indexName] ON mytable; 


-- 唯一索引
-- 它与前面的普通索引类似，不同的就是：索引列的值必须唯一，但允许有空值。如果是组合索引，则列值的组合必须唯一。
CREATE UNIQUE INDEX indexName ON mytable(username(length)) 


显示索引信息
你可以使用 SHOW INDEX 命令来列出表中的相关的索引信息。可以通过添加 \G 来格式化输出信息。
mysql> SHOW INDEX FROM table_name\G