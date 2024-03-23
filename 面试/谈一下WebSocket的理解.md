



Websocket 握手

```
GET /chat HTTP/1.1
Host: server.example.com
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: x3JJHMbDL1EzLkh9GBhXDw==
Sec-WebSocket-Protocol: chat, superchat
Sec-WebSocket-Version: 13
Origin: http://example.com
```


然后服务器会返回下列东西，表示已经接受到请求， 成功建立 WebSocket 

```
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: HSmrc0sMlYUkAGmm5OPpG2HaGWk=
Sec-WebSocket-Protocol: chat
```


用express 写一个websocket 服务端

```
npm install ws
```


```js
const express = require('express');
const http = require('http');
const WebSocket = require('ws');

const PORT = 3000;
const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

wss.on('connection', (ws) => {
    console.log('Client connected');
    const interval = setInterval(() => {
        ws.send('hello world');
    }, 1000);

    ws.on('close', () => {
        console.log('Client disconnected');
        clearInterval(interval);
    });
});

server.listen(PORT, () => {
    console.log(`WebSocket server running on ws://localhost:${PORT}/websocket/socketTest`);
});
```



写一个websocket的前端接受服务器信息

```js
import './App.css'
import React, { useState, useEffect } from 'react';

function App() {
    const [response, setResponse] = useState('');

    useEffect(() => {
        const socket = new WebSocket('ws://localhost:3000/websocket/socketTest');

        socket.onopen = () => {
          setResponse('WebSocket connection opened');
        };

        socket.onmessage = (event) => {
            console.log(event);
            setResponse(event.data);
        };

        return () => {
          if (socket.readyState === 1) { // <-- This is important
              socket.close();
          }
      }
    }, []);

    return (
        <div>
            <p>Response from WebSocket: {response}</p>
        </div>
    );
}

export default App;

```