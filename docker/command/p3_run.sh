docker run  -v $PWD/run_code/python:/usr/src/myapp  -w /usr/src/myapp python:3.5 python helloworld.py
docker run --name python-test python:3.5
docker run --name python-test -it python:3.5 bash

docker run -itd --name node-test node



docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

docker run -d \
    --name some-postgres \
    -e POSTGRES_PASSWORD=mysecretpassword \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /custom/mount:/var/lib/postgresql/data \
    postgres

docker run -d \
    --name my-postgres \
    -e POSTGRES_PASSWORD=mysecretpassword \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    postgres

docker run -p 8080:8080 -p 9990:9990 -it jboss/wildfly /opt/jboss/wildfly/bin/standalone.sh -b 0.0.0.0 -bmanagement 0.0.0.0

docker run -d -p 8888:8080 tomcat

-p host:container

docker run -itd --name ubuntu-test ubuntu


