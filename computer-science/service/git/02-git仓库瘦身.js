/**
 *
 *
 * https://rtyley.github.io/bfg-repo-cleaner/
 * 
根据需要，运行 bfg 命令，例如清理所有超过 10M 的文件：
bfg --strip-blobs-bigger-than 10M your-repo.git

清理前 100 大的文件：
bfg --strip-biggest-blobs 100 your-repo.git


指定文件或按照通配符清理文件：
bfg --delete-files "node.exe" your-repo.git
bfg --delete-files "*.zip" your-repo.git


运行下面的命令对仓库历史进行改写，运行速度会比较慢：
cd your-repo.git
git reflog expire --expire=now --all && git gc --prune=now --aggressive
$ git push

 */
