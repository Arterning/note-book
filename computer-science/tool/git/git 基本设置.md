## git 设置用户名密码

```tsx
$ git config --global user.name Arterning
$ git config --global user.email ning.huang.ridson@gmail.com
```

## git配置文件

cat ~/.gitconfig

```bash
[user]
    name = gitaccount
    email = gitaccount@example.com
```

```bash
$ git config --global --unset user.name
$ git config --global --unset user.email
$ git config --global user.name
#全局配置账户已经移除
$ git config --global user.email
#全局配置邮箱已经移除
```

```tsx
$ cd git-repository/
$ git config user.name anothergitaccount
$ git config user.email anothergitaccount@example.com
```

## git 修改提交人姓名

```bash
git filter-branch --env-filter '

OLD_EMAIL="hnbingo@petalmail.com"
CORRECT_NAME="Arterning"
CORRECT_EMAIL="ning.huang.ridson@gmail.com"

if [ "$GIT_COMMITTER_EMAIL" = "$OLD_EMAIL" ]
then
export GIT_COMMITTER_NAME="$CORRECT_NAME"
export GIT_COMMITTER_EMAIL="$CORRECT_EMAIL"
fi
if [ "$GIT_AUTHOR_EMAIL" = "$OLD_EMAIL" ]
then
export GIT_AUTHOR_NAME="$CORRECT_NAME"
export GIT_AUTHOR_EMAIL="$CORRECT_EMAIL"
fi
export GIT_AUTHOR_NAME="$CORRECT_NAME"
export GIT_AUTHOR_EMAIL="$CORRECT_EMAIL"
' --tag-name-filter cat -- --branches --tags

git filter-branch --env-filter '

CORRECT_NAME="Arterning"
CORRECT_EMAIL="ning.huang.ridson@gmail.com"
export GIT_AUTHOR_NAME="$CORRECT_NAME"
export GIT_AUTHOR_EMAIL="$CORRECT_EMAIL"
' --tag-name-filter cat -- --branches --tags
```

```bash
##git 修改commit 时间
git commit --date="May 5 13:05:20 2023 +0800" -am "Rename dir"
```