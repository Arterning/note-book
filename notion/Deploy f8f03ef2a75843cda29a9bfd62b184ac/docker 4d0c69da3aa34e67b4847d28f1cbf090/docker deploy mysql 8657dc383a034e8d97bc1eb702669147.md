# docker deploy mysql

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