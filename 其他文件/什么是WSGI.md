
wiki 上的解释

- Web Server Gateway Interface
- Web 服务器网关接口
- 是为 Python 定义的 Web 服务器和 Web 应用程序或框架之间的一种简单而通用的接口


#### WSGI 的问题

- 随着移动网络的发展，Web 技术也在升级，比如 WebSocket、HTTP/2，HTTP/3
- WSGI 应用是一个单调用、同步接口，即输入一个请求，返回一个响应；这个模式无法支持长连接或者 WebSocket 这样的连接
- 即使想办法将 WSGI 应用改成异步，还有另一个限制：一个 URL 对应一个请求，而 HTTP/2、Websocket 等在一个 URL 里会出现多个请求
