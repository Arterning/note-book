sqlmap，怎么对⼀个注⼊点注⼊？
1）如果是 get 型号，直接，sqlmap -u "诸如点⽹址".
2) 如果是 post 型诸如点，可以 sqlmap -u "注⼊点⽹址” --data="userid=aaa&passwd=bbbb"
3）如果是 cookie，X-Forwarded-For 等，可以访问的时候，⽤ burpsuite 抓包，注⼊处⽤号替换，放到⽂件⾥，然后sqlmap -r "⽂件地址

发现 demo.jsp?uid=110 注⼊点，你有哪⼏种思路获取 webshell，哪种是优选？
有写⼊权限的，构造联合查询语句使⽤ using INTO OUTFILE，可以将查询的输出重定向到系统的⽂件中，这样去写⼊WebShell 使⽤ sqlmap –os-shell 原理和上⾯⼀种相同，来直接获得⼀个 Shell，这样效率更⾼ 通过构造联合查询语句得到⽹站管理员的账户和密码，然后扫后台登录后台，再在后台通过改包上传等⽅法上传 Shell

以下链接存在 sql 注⼊漏洞，对于这个变形注⼊，你有什么思路？
demo.do?DATA=AjAxNg==
DATA 有可能经过了 base64 编码再传⼊服务器，所以我们也要对参数进⾏ base64 编码才能正确完成测试