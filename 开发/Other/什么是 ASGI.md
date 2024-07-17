## 什么是 ASGI 服务器？

答： 异步网关协议接口，一个介于网络协议服务和 Python 应用之间的标准接口，能够处理多种通用的协议类型，包括 HTTP，HTTP2 和 WebSocket。

Uvicorn 就是一种ASGI服务器


## 什么是 Uvicorn ？

答：Uvicorn 是基于 uvloop 和 httptools 构建的非常快速的 ASGI 服务器。相当于Nodejs中的Express/Fastify


## 什么是 uvloop 和 httptools ？

答： uvloop 用于替换标准库 asyncio 中的事件循环，使用 Cython 实现，它非常快，可以使 asyncio 的速度提高 2-4 倍。asyncio 不用我介绍吧，写异步代码离不开它。

httptools 是 nodejs HTTP 解析器的 Python 实现。







## 使用方法：

```python
$ pip install uvicorn
```

创建一个文件 example.py

```python
async def app(scope, receive, send):
    assert scope['type'] == 'http'
    await send({
        'type': 'http.response.start',
        'status': 200,
        'headers': [
            [b'content-type', b'text/plain'],
        ]
    })
    await send({
        'type': 'http.response.body',
        'body': b'Hello, world!',
    })
```

启动 Uvicorn

```python
$ uvicorn example:app
```

你也可以不使用命令行，直接运行你的脚本也是可以的，如下：

```python
import uvicorn

async def app(scope, receive, send):
    ...

if __name__ == "__main__":
    uvicorn.run("example:app", host="127.0.0.1", port=5000, log_level="info")
```

使用命令行时，你可以使用 `uvicorn --help` 来获取帮助。