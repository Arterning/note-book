# docker-compose-volumes的说明
docker-compose中有两种方式可以设置数据持久化。
1. 绝对路径
```yml
db:
    image: mariadb:latest
    restart: always
    ports:
      - "3306:3306"
    volumes:
      - [直接使用宿主机的本地路径]:/var/lib/mysql

```

2. 卷标

```yml
db:
    image: mariadb:latest
    restart: always
    ports:
      - "3306:3306"
    volumes:
      - [卷标名称]:/var/lib/mysql

volumes:
  [卷标名称]:
```

使用卷标的方式，简洁易于管理，但是数据实际存放的位置需要费点周折才能看到。
查看卷标的信息

```bash
#输出的信息中 Mountpoint 即为数据实际存放的位置
docker volume inspect [卷标名称]


#还可以查看所有的卷标
docker volume ls
```