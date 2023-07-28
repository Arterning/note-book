# Fisrt use AI !!!! What a impressive feature!!!!!

# How to Deploy MySQL by Docker Compose?

Docker Compose is a tool for defining and running multi-container Docker applications. It simplifies the process of managing multiple containers by allowing you to define them in a single YAML file and run them together with a single command. In this guide, we will walk through the steps to deploy MySQL by Docker Compose.

## Prerequisites

Before we begin, make sure you have the following installed on your machine:

- Docker (latest version)
- Docker Compose (latest version)

## Step 1: Create a Docker Compose file

The first step is to create a new directory where we will create a `docker-compose.yml` file. You can name the directory whatever you like. Inside the directory, create a new file called `docker-compose.yml` and add the following code to it:

```
version: '3'

services:
  mysql:
    image: mysql:latest
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: your_password
    ports:
      - "3306:3306"
    volumes:
      - ./mysql-data:/var/lib/mysql

```

This code defines a service called `mysql` using the latest MySQL image. It also sets the root password for the MySQL server and maps the port `3306` to the host machine so that you can connect to the MySQL server from outside the container. Finally, it mounts the `./mysql-data` directory on the host machine to the `/var/lib/mysql` directory in the container to persist the MySQL data.

## Step 2: Start the MySQL server

Now that we have defined our service in the `docker-compose.yml` file, we can start the MySQL server by running the following command in the directory where the `docker-compose.yml` file is located:

```
docker-compose up -d

```

This command will download the MySQL image (if it is not already downloaded) and start the container in the background. The `-d` flag tells Docker Compose to run the containers in detached mode, which means that the containers will run in the background and not print their output to the console.

## Step 3: Verify the MySQL server is running

To verify that the MySQL server is running, run the following command:

```
docker ps

```

This command will display a list of running containers. You should see a container with the name `docker_mysql_1` (assuming the directory you created is called `docker`). This container is running the MySQL server.

## Conclusion

In this guide, we walked through the steps to deploy MySQL by Docker Compose. Docker Compose makes it easy to define and run multi-container Docker applications, and is an essential tool for anyone working with Docker. By following the steps outlined in this guide, you should now have a working MySQL server running in a Docker container on your machine.

## Step 4: Connect to the MySQL server

Now that the MySQL server is running, you can connect to it using a MySQL client. There are many MySQL clients available, but one of the most popular is the MySQL command-line client. To install the MySQL client, you can run the following command:

```
sudo apt-get install mysql-client

```

Once the client is installed, you can connect to the MySQL server by running the following command:

```
mysql -h localhost -u root -p

```

This command connects to the MySQL server running on the local machine (`localhost`) using the root user and prompts you for the root password. Enter the password you set in the `docker-compose.yml` file and press Enter.

If everything is working correctly, you should now be connected to the MySQL server and see a prompt that looks like this:

```
mysql>

```

## Step 5: Create a database

Now that we are connected to the MySQL server, we can create a new database. To create a new database, run the following command:

```
CREATE DATABASE mydatabase;

```

This command creates a new database called `mydatabase`. You can replace `mydatabase` with any name you like.

## Step 6: Verify the database is created

To verify that the database was created, run the following command:

```
SHOW DATABASES;

```

This command will display a list of all the databases on the MySQL server. You should see the `mydatabase` database in the list.

## Conclusion

In this guide, we walked through the steps to deploy MySQL by Docker Compose and connect to it using a MySQL client. We also created a new database and verified that it was created successfully. Docker Compose makes it easy to manage multi-container Docker applications, and is an essential tool for anyone working with Docker.

To persist the MySQL data, you need to mount a volume from the host machine to the container. In the `docker-compose.yml` file, add a `volumes` section to the `mysql` service definition as follows:

```
volumes:
  - ./mysql-data:/var/lib/mysql

```

This code mounts the `./mysql-data` directory on the host machine to the `/var/lib/mysql` directory in the container. This will allow the MySQL data to persist even if the container is destroyed or recreated.