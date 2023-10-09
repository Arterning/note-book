```bash
## 运行 bfg 命令，例如清理所有超过 10M 的文件
java -jar .\\bfg-1.14.0.jar --strip-blobs-bigger-than 2M .\\web-dev-examples\\
java -jar .\\bfg-1.14.0.jar --strip-biggest-blobs 100  .\\web-dev-examples\\
```

按照通配符清理文件

```bash
bfg --delete-files "*.jpg" your-repo.git
bfg --delete-files "*.png" your-repo.git
bfg --delete-files "*.zip" your-repo.git
```

```bash
git reflog expire --expire=now --all && git gc --prune=now --aggressive
git reflog expire --expire=now --all ; git gc --prune=now --aggressive
git push origin --force
```

- [bfg repo 官网](https://rtyley.github.io/bfg-repo-cleaner/)