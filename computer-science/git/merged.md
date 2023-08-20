# 换行符配置

core.autocrf
- true 提交时改成LF，检出时改成CRLF:
- input:提交时改成LF，检出时不改
- false: 提交时是什么就是什么，不改换行符，检出时也不改

core.safecrf
- true: 拒绝提交包含混合换行符的文件
- false: 允许提交包含混合换行符的文件
- warn: 提交包含混合换行符的文件时给出警告
(会提示 Fatal:xxx)


# 00-git配置信息管理.md


 config 配置有system级别 global（用户级别） 和local（当前仓库）三个 设置先从system-》global-》local  底层配置会覆盖顶层配置 分别使用--system/global/local 可以定位到配置文件

 * 查看系统config
   git config --system --list

   查看当前用户（global）配置
   git config --global  --list

   查看当前仓库配置信息
   git config --local  --list
 */


# 01-git设置代理.md

/**
 * 

使用http代理 
git config --global http.proxy http://127.0.0.1:58591
git config --global https.proxy https://127.0.0.1:58591
#使用socks5代理
git config --global http.proxy socks5://127.0.0.1:51837
git config --global https.proxy socks5://127.0.0.1:51837

关于http代理和socks代理的区别这里稍微提一句

1 https 协议只支持 http/https，一般的 ie 代理用的 http/https 协议。如果是应用层协议一般不用 http/https，有些应用程序只能使用 socks 代理。

2 socks 工作在会话层上，而 http 工作在应用层上，socks 代理只是简单地传递数据包，而不必关心是何种应用协议(比如 FTP、HTTP 和 NNTP 请求)，所以 socks 代理服务器比应用层代理服务器要快得多。

需要注意的是 socks协议也可以代理Http请求 而且更快 也就是git的这种设置方式
git config --global http.proxy socks5://127.0.0.1:51837
git config --global https.proxy socks5://127.0.0.1:51837

只对Github代理（推荐）
#使用socks5代理（推荐）
git config --global https.https://github.com.proxy socks5://127.0.0.1:51837
#使用http代理（不推荐）
git config --global https.https://github.com.proxy http://127.0.0.1:58591

取消代理
当你不需要使用代理时，可以取消之前设置的代理。

git config --global --unset http.proxy
git config --global --unset https.proxy


对于windows用户，代理会用到connect.exe
设置ssh代理（终极解决方案）
 ProxyCommand "C:\Program Files\Git\mingw64\bin\connect" -S 127.0.0.1:10808 -a none %h %p
 */


# 02-git仓库瘦身.md

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


# 03-git修改提交时间.md

/**

Thanks to ChatGPT

Set the environment variables:

$ export GIT_AUTHOR_DATE="<desired date>"
$ export GIT_COMMITTER_DATE="<desired date>"
Replace "<desired date>" with the desired date and time, in the format "YYYY-MM-DD HH:MM:SS".

export GIT_COMMITTER_DATE="2023-01-26 08:50:55"
export GIT_AUTHOR_DATE="2023-01-19 08:50:55"

Create the commit:
$ git commit -m "<commit message>"
 */

# 04-git_branch.md

/**
 *  
 创建分支命令：
 git branch (branchname)

 git checkout (branchname)
 git branch


 创建新分支并立即切换到该分支
 git checkout -b newtest
 git branch -d (branchname)
 */


# 05-git_push.md

# git push --set-upstream gitee master
git add .
git commit 'commit'
git push github
git push gitee

# 08-git-autocrf.md

/**
 * 今天在使用git add 命令的时候，弹出了一个警告 warning: LF will be replaced by CRLF in ******（具体的一个文件）

原因
LF和CRLF其实都是换行符，但是不同的是，LF是linux和Unix系统的换行符，CRLF是window 系统的换行符。这就给跨平台的协作的项目带来了问题，保存文件到底是使用哪个标准呢？ git为了解决这个问题，提供了一个”换行符自动转换“的功能，并且这个功能是默认处于”自动模式“即开启状态的。
这个换行符自动转换会把自动把你代码里 与你当前操作系统不相同的换行的方式 转换成当前系统的换行方式（即LF和CRLF 之间的转换），这样一来，当你提交代码的时候，即使你没有修改过某个文件，也被git认为你修改过了，从而提示"LF will be replaced by CRLF in *****"

解决
最简单的一种办法就是把自动转换功能关掉即可。
输入命令 ：git config core.autocrlf false (仅对当前git仓库有效）
git config --global core.autocrlf false (全局有效，不设置推荐全局）

然后重新提交代码即可。

================ 2021年3月30日15点14分 更新 ============
这篇博客大概是大二的时候写的了，当时对git也是初步接触，现在回过头来看，其实这个warning无伤大雅。在实际编程中，一般来说warning级别的警告都是可以忽略的，这个警告也是。
如果你是一个强迫症，非要去掉所有warning，也是可以。如果你可以保证你的代码不会跨平台开发，（比如你和你的合作者用不同的系统进行开发时，关掉这个自动转换的功能可能会导致代码显示异常），你可以设置关掉自动转换的功能。
当然，结合实际情况来说，你不能保证你的所有代码都不会跨平台开发，因为你不能保证你的合作者用的是跟你一样的系统，这个时候，最好就是只针对当前仓库设置，你只要保证当前仓库的代码不会跨平台开发就行。
就我现在来说，比较建议忽略这个警告。
 * 
 * 
 */




