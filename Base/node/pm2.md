
```bash

pm2 start app.js --name app
pm2 list
pm2 ls
pm2 delete app
pm2 stop app
pm2 restart app
pm2 info app
pm2 logs app

# 监控(monitor)-查看进程的资源消耗情况
pm2 monit

```


pm2 start
-i --instances：启用多少个实例，可用于负载均衡。如果-i 0或者-i max，则根据当前机器核数确定实例数目。

PM2 can automatically restart your application when a file is modified in the current directory or its subdirectories:

```
pm2 start app.js --watch
```



Run npm script
```
pm2 start 'bun run start'
```


## pm2 开启前端服务

```
pm2 start pnpm -- run dev --host 0.0.0.0
pm2 list
pm2 restart 0
```