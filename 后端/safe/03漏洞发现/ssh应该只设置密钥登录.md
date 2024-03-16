ssh应该只设置密钥登录

如果使用密码登录 就会报错：


>
$ ssh root@141.164.57.221
The authenticity of host '141.164.57.221 (141.164.57.221)' can't be established.
ED25519 key fingerprint is SHA256:VaEy8zmFHQ7jELgSqyYAbMRB44sHtacFtEvhN4DSMsY.
This key is not known by any other names
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '141.164.57.221' (ED25519) to the list of known hosts.
root@141.164.57.221: Permission denied (publickey).

00775626@Z-A61-ZK0040 MINGW64 ~
