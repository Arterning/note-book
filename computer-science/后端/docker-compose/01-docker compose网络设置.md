# Docker Compose 网络设置

# 基本概念
举个例子，假如一个应用程序在名为myapp的目录中，并且docker-compose.yml如下所示：
```yml
version: '2'
services:
  web:
    build: .
    ports:
      - "8000:8000"
  db:
    image: postgres

```
当我们运行docker-compose up时，将会执行以下几步：
- 创建一个名为myapp_default的网络；
- 使用web服务的配置创建容器，它以“web”这个名称加入网络myapp_default；
- 使用db服务的配置创建容器，它以“db”这个名称加入网络myapp_default。

容器间可使用服务名称（web或db）作为hostname相互访问。例如，web这个服务可使用postgres://db:5432 访问db容器。

# 更新容器
当服务的配置发生更改时，可使用docker-compose up命令更新配置。
此时，Compose会删除旧容器并创建新容器。新容器会以不同的IP地址加入网络，名称保持不变。任何指向旧容器的连接都会被关闭，容器会重新找到新容器并连接上去。

# Link
前文讲过，默认情况下，服务之间可使用服务名称相互访问。links允许我们定义一个别名，从而使用该别名访问其他服务。举个例子：
```yml
version: '2'
services:
  web:
    build: .
    links:
      - "db:database"
  db:
    image: postgres
```
该选项是 docker 历史遗留的选项, 目前已被用户自定义网络名称空间取代, 最终有可能被废弃 (在使用 swarm 部署时将忽略该选项)


# 指定自定义网络
一些场景下，默认的网络配置满足不了我们的需求，此时我们可使用networks命令自定义网络。networks命令允许我们创建更加复杂的网络拓扑并指定自定义网络驱动和选项。不仅如此，我们还可使用networks将服务连接到不是由Compose管理的、外部创建的网络。如下，我们在其中定义了两个自定义网络。

```yml
version: '2'

services:
  proxy:
    build: ./proxy
    networks:
      - front
  app:
    build: ./app
    networks:
      - front
      - back
  db:
    image: postgres
    networks:
      - back

networks:
  front:
    # Use a custom driver
    driver: custom-driver-1
  back:
    # Use a custom driver which takes special options
    driver: custom-driver-2
    driver_opts:
      foo: "1"
      bar: "2"
```
其中，proxy服务与db服务隔离，两者分别使用自己的网络；app服务可与两者通信。
由本例不难发现，使用networks命令，即可方便实现服务间的网络隔离与连接。



# 使用已存在的网络
一些场景下，我们并不需要创建新的网络，而只需加入已存在的网络，此时可使用external选项。示例：
```yml
networks:
  default:
    external:
      name: my-pre-existing-network
```

# Docker Compose 链接外部容器
在Docker中，容器之间的链接是一种很常见的操作：它提供了访问其中的某个容器的网络服务而不需要将所需的端口暴露给Docker Host主机的功能。

在不使用Docker Compose的时候，将两个容器链接起来使用—link参数，相对来说比较简单，以nginx镜像为例子：

```yml
docker run --rm --name test1 -d nginx  #开启一个实例test1
docker run --rm --name test2 --link test1 -d nginx #开启一个实例test2并与test1建立链接
```

这样，test2与test1便建立了链接，就可以在test2中使用访问test1中的服务了。
如果使用Docker Compose，那么这个事情就更简单了

```yml
version: "3"
services:
  test2:
    image: nginx
    depends_on:
      - test1
    links:
      - test1
  test1:
    image: nginx
```
最终效果与使用普通的Docker命令docker run xxxx建立的链接并无区别。这只是一种最为理想的情况。



# 例外的情况
如果容器没有定义在同一个docker-compose.yml文件中，应该如何链接它们呢？
又如果定义在docker-compose.yml文件中的容器需要与docker run xxx启动的容器链接，需要如何处理？

1. 方式一：让需要链接的容器同属一个外部网络
我们还是使用nginx镜像来模拟这样的一个情景：假设我们需要将两个使用Docker Compose管理的nignx容器（test1和test2）链接起来，使得test2能够访问test1中提供的服务，这里我们以能ping通为准。首先，我们定义容器test1

```yml
version: "3"
services:
  test2:
    image: nginx
    container_name: test1
    networks:
      - default
      - app_net
networks:
  app_net:
    external: true
```
容器test2内容与test1基本一样，只是多了一个external_links,需要特别说明的是：最近发布的Docker版本已经不需要使用external_links来链接容器，容器的DNS服务可以正确的作出判断，因此如果你你需要兼容较老版本的Docker的话，那么容器test2的docker-compose.yml文件内容为

```yml
version: "3"
services:
  test2:
    image: nginx
    networks:
      - default
      - app_net
    external_links:
      - test1
    container_name: test2
networks:
  app_net:
    external: true
```

正如你看到的那样，这里两个容器的定义里都使用了同一个外部网络app_net,因此，我们需要在启动这两个容器之前通过以下命令再创建外部网络：

```shell
docker network create app_net
```

如果我们通过docker run --rm --name test3 -d nginx这种方式来先启动了一个容器(test3)并且没有指定它所属的外部网络，而需要将其与test1或者test2链接的话，这个时候手动链接外部网络即可：

```shell
docker network connect app_net test3
```
这样，三个容器都可以相互访问了。


2. 方式二：更改需要链接的容器的网络模式
通过更改你想要相互链接的容器的网络模式为bridge,并指定需要链接的外部容器（external_links)即可。与同属外部网络的容器可以相互访问的链接方式一不同，这种方式的访问是单向的。

还是以nginx容器镜像为例子，如果容器实例nginx1需要访问容器实例nginx2，那么nginx2的doker-compose.yml定义为：

```yml
version: "3"
services:
  nginx2:
    image: nginx
    container_name: nginx2
    network_mode: bridge
```

nginx1的docker-compose.yml定义为：


```yml
version: "3"
services:
  nginx1:
    image: nginx
    external_links:
      - nginx2
    container_name: nginx1
    network_mode: bridge

```

接着我们使用ping来测试下连通性：

```bash
$ docker exec -it nginx1 ping nginx2  # nginx1 to nginx2
PING nginx2 (172.17.0.4): 56 data bytes
64 bytes from 172.17.0.4: icmp_seq=0 ttl=64 time=0.141 ms
64 bytes from 172.17.0.4: icmp_seq=1 ttl=64 time=0.139 ms
64 bytes from 172.17.0.4: icmp_seq=2 ttl=64 time=0.145 ms

$ docker exec -it nginx2 ping nginx1 #nginx2 to nginx1
ping: unknown host

```

# 什么是external networks
有时候我们想使用我们通过docker network create创建好的网络，而不是让docker-compose创建一个新的，这个时候就需要用到“external”关键字了。表示在外部用network create创建的网络


