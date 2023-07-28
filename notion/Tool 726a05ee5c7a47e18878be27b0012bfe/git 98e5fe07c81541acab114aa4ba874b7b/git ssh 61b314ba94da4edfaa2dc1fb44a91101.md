# git ssh

## 密钥配置

首先生成密钥

```bash
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

接着在你的服务器或者git server 添加你的公钥，私钥在你本地保存

## 密钥登陆

当使用 ssh -T git@github.com 登录之后

在.ssh 目录下会创建config文件夹

在linux平台下的配置

```bash
Host gitee.com
        HostName gitee.com
        User git
        IdentityFile .ssh/id_rsa.gitee
Host github.com
        HostName github.com
        User git
        IdentityFile .ssh/id_rsa.github
Host gitlab.com
        HostName gitlab.com
        User git
        IdentityFile .ssh/id_gitlab_rsa
```

在windows下的配置

```bash
Host gitee.com
        HostName gitee.com
        User git
        IdentityFile C:\Users\china\.ssh\id_rsa.gitee
Host github.com
        HostName github.com
        User git
        IdentityFile C:\Users\china\.ssh\id_rsa.github
Host gitlab.com
        HostName gitlab.com
        User git
        IdentityFile C:\Users\china\.ssh\id_gitlab_rsa
```

## 密钥权限

如果登录出现报错

```bash
no such identity: .ssh/id_rsa.github: No such file or directory
git@github.com: Permission denied (publickey).
```

```bash
$ sudo chmod 600 ~/.ssh/id_rsa
$ sudo chmod 644 ~/.ssh/id_rsa.pub
```