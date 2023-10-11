_QPS_的概念如下所示: _QPS_（Query Per Second）：每秒请求数，就是说服务器在一秒的时间内处理了多少个请求

## redis的QPS

运行 Redis 自带的基准测试工具，运行 set，get 1000000 次，1s 后退出并显示数据

`Redis` 的大致数据为 50000 - 300000

此时一句话脱口而出重剑无锋，大巧不工，Redis 的优秀性能真的不是擂的，正如 《Redis in Action》中提到，Redis 的性能会是普通关系型数据库的 10 - 100 倍

## mysql的QPS

MySQL 的 QPS 为 4000 左右浮动

## 总结

基准测试就可以用来测试QPS