-u 单个URL -m xx.txt 多个URL
-d "mysql://user:password@10.10.10.137:3306/dvwa"  作为服务器客户端，直接连接数据库
--data post/get都适用
-p 指定扫描的参数
-r 读取文件
-f 指纹信息
--tamper 混淆脚本，用于应用层过滤
--cookie --user-agent --host等等http头的修改
--threads 并发线程 默认为1
--dbms MySQL<5.0> 指定数据库或版本

–level=LEVEL 执行测试的等级（1-5，默认为 1）
–risk=RISK 执行测试的风险（0-3，默认为 1） Risk升高可造成数据被篡改等风险
–current-db / 获取当前数据库名称
–dbs 枚举数据库管理系统数据库
–tables 枚举 DBMS 数据库中的表
–columns 枚举 DBMS 数据库表列

-D DB 要进行枚举的数据库名
-T TBL 要进行枚举的数据库表
-C COL 要进行枚举的数据库列
-U USER 用来进行枚举的数据库用户