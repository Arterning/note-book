# git-merge-repo

# git 合并多个仓库

基本的步骤：

- 在主 Git 仓库中添加要合并的远程 Git 仓库作为远程仓库。

```bash
git remote add <name1> <url-to-repo1>git remote add <name2> <url-to-repo2>
```

- 从每个远程仓库中获取最新的提交记录

```bash
git fetch <name1>git fetch <name2>
```

- 将每个远程仓库的主分支合并到主 Git 仓库的主分支中。

```bash
git merge <name1>/<branch> --allow-unrelated-historiesgit merge <name2>/<branch> --allow-unrelated-histories
```

- `-allow-unrelated-histories` 是一个 Git 命令选项，它在使用 git merge 命令合并两个没有共同历史记录的分支时起作用。通常情况下，Git 不允许这样的合并，因为两个分支没有共同的祖先，Git 无法确定如何合并它们的历史记录。

使用 `--allow-unrelated-histories` 选项可以告诉 Git，让它忽略两个分支没有共同历史记录的限制，并尝试将它们合并起来。这样，就可以将两个独立的 Git 仓库合并为一个仓库，而不会丢失任何历史记录。

需要注意的是，在使用 `--allow-unrelated-histories` 选项合并分支时，可能会产生冲突，需要手动解决这些冲突。此外，合并后的仓库可能会变得非常大，需要谨慎使用。