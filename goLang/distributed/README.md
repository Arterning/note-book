<!--
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 14:05:17
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 16:01:30
 * @FilePath: \code\goLang\new\README.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
分布式模型有三种
1. Hub&Spoke
2. Peer to Peer
3. Message Queue


创建Web服务
创建注册服务
注册Web服务
取消注册Web服务


curl -H "Content-Type: application/json" -X POST -d 'fuck you' \
"http://127.0.0.1:4000/log"


curl -H "Content-Type: application/json" -X POST -d '{"ServiceName":"abc", "ServiceURL":"github.com"}' \
"http://127.0.0.1:3000/services"