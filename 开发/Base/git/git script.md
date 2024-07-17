写一个git脚本 接受2个参数source 和 target 先更新source，然后把target合并到source，最后push source


```
#!/bin/bash

# 获取参数
source_branch="$1"
target_branch="$2"


git fetch -p

# 更新source分支
git checkout "$source_branch"
git pull

# 合并target分支到source分支
git merge "$target_branch"

# 推送source分支
git push origin "$source_branch"
```


在`git fetch`命令中，`-p`选项表示`--prune`，它的作用是在从远程仓库获取分支更新时，同时删除本地仓库中已经不存在于远程仓库的分支。

具体来说，当你运行`git fetch -p`时，Git会从远程仓库获取最新的分支信息，并将这些更新同步到本地仓库。如果在本地仓库中存在一些分支，但在远程仓库中已经被删除了，那么使用`-p`选项会自动删除这些本地分支，以保持本地仓库与远程仓库的一致性。

这个选项在清理本地仓库中已经被删除的分支时非常有用，以确保你的本地仓库与远程仓库保持同步。

