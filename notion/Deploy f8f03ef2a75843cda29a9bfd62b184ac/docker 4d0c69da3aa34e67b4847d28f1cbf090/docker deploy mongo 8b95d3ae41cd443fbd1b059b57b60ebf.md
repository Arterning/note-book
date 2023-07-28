# docker deploy mongo

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