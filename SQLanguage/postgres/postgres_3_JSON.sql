-- postgresql支持两种json数据类型：json和jsonb，
-- 而两者唯一的区别在于效率,json是对输入的完整拷贝，使用时再去解析，所以它会保留输入的空格，重复键以及顺序等。
-- 而jsonb是解析输入后保存的二进制，它在解析时会删除不必要的空格和重复的键，顺序和输入可能也不相同。使用时不用再次解析。两
-- 者对重复键的处理都是保留最后一个键值对。效率的差别：json类型存储快，使用慢，jsonb类型存储稍慢，使用较快。
--
-- 注意：键值对的键必须使用双引号

drop table json_table;
create table json_table
(
    ID   INT PRIMARY KEY NOT NULL,
    NAME TEXT,
    jjdoc json,
    jdoc jsonb
);

insert into json_table (id, jdoc)
values (123, '{
  "guid": "9c36adc1-7fb5-4d5b-83b4-90356a46061a",
  "name": "Angela Barton",
  "is_active": true,
  "company": "Magnafone",
  "address": "178 Howard Place, Gulf, Washington, 702",
  "registered": "2009-11-07T08:53:22 +08:00",
  "latitude": 19.793713,
  "longitude": 86.513373,
  "tags": [
    "enim",
    "aliquip",
    "qui"
  ]
}');

select *from json_table;



SELECT '{"bar": "baz", "balance":      7.77, "active":false}'::json;
select row_to_json(row(1,'foo',3.23,'dkdd'));


--XML
show xmloption;
SET xmloption TO document;

select xmlparse(document'<title> hello world</title>');