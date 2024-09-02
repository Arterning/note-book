

```java
version: '3'

services:
  mysql:
    image: mysql:8
    restart: always
    environment:
      MYSQL_DATABASE: wipi
      MYSQL_ROOT_PASSWORD: root
    ports:
      - "3306:3306"
    volumes:
      - ./mysql-data:/var/lib/mysql
```



使用的配置文件

```yaml
version: '3'

services:
  mysql:
    image: mysql:8
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: L94K5qLmEUk4xDtJfTu1
    ports:
      - "3306:3306"
    volumes:
      - ./mysql-data:/var/lib/mysql

  adminer:
    image: adminer
    container_name: mysql-adminer
    restart: always
    ports:
      - 8080:8080

  redis:
    image: redis
    command: redis-server --requirepass 1568456452
    ports:
    - "6379:6379"

  nginx:
    image: nginx
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx-data/nginx.conf:/etc/nginx/nginx.conf
      - ./nginx-data/conf.d:/etc/nginx/conf.d
      - ./nginx-data/ssl:/etc/nginx/ssl
      - ./nginx-data/666:/usr/share/nginx/html/666
```

pgsql配置

```yaml
version: '3'

services:
  db:
    image:  postgres
    restart: always
    ports:
      - "5432:5432"
    environment:
      POSTGRES_PASSWORD: oPOdLbEPvGJJXB9Myaf6
      POSTGRES_DB: test
  test-db:
    image: postgres
    restart: always
    ports:
      - "5433:5432" # 👈 Note the 5433 port (since we are using 5432 for our regular db)
    environment:
      POSTGRES_PASSWORD: oPOdLbEPvGJJXB9Myaf6
```


## RabbitMQ
```jsx
version: "3.2"
services:
  rabbitmq:
    image: rabbitmq:3-management-alpine
    container_name: 'rabbitmq'
    ports:
        - 5672:5672
        - 15672:15672
    volumes:
        - ~/.docker-conf/rabbitmq/data/:/var/lib/rabbitmq/
        - ~/.docker-conf/rabbitmq/log/:/var/log/rabbitmq
    networks:
        - rabbitmq_go_net

networks:
  rabbitmq_go_net:
    driver: bridge
```

默认登录 guest/guest

go to localhost:15672

## Mongo
```jsx
version: '3.1'
services:
    db:
        image: mongo
        restart: always
        container_name: mongodb-container
        ports:
            - '27017:27017'
        environment:
          MONGO_INITDB_ROOT_USERNAME: root
          MONGO_INITDB_ROOT_PASSWORD: xiaohuiwoaini12345
        volumes:
            - ./mongo-data:/data/db
```