# drone教程
Drone CI for Github 的部署到此就真的结束了，一路走来，踩了不少坑，Drone 的文档你慢慢看了以后就会发现有多烂，烂到心累想哭，许多在 Demo 中出现的变量在参考手册中找不到，完全不知道什么含义，只能靠瞎猜和摸索。

有的东西文档烂 真是不好学啊 包括开源代码也是 如果文档烂
根本不好学

- https://learnku.com/articles/50293
- https://www.drone.io/


# 打造自己的jenkis ci
有的项目 本身技术特点比较负责 你要部署起来就比较麻烦
一般来说 
1. 你要下载依赖 maven jar包 node js包 经常出现下载缓慢 下载失败的情况
2. 你要编译二进制文件
3. 你要配置数据库。。
所以把这一整套流程用Dockerfile描述 用镜像承载这一切是非常好的办法
同时结合CI 当代码推送到github的时候自动触发构建 更新环境。

# vaultwarden项目学习
1. vaultwarden 前端Docker仓库 https://hub.docker.com/r/vaultwarden/web-vault/tags

2. 对应的移动端 使用的是C# Xamarin构建的 https://github.com/bitwarden/mobile

3. xamarin官网 https://dotnet.microsoft.com/zh-cn/apps/xamarin

4. 用xamarin开发App的体验如何？ https://www.zhihu.com/question/31159902


# other
1. bitwarden官网 https://bitwarden.com/

2. 为什么国内的uniapp一直没人讨论呢？ https://www.zhihu.com/question/330607558
