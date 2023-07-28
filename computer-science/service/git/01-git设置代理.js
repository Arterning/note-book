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
