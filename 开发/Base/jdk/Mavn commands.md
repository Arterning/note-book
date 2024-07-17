
## start up spring-boot application via command line

I presume you are trying to compile the application and run it without using an IDE. I also presume you have maven installed and correctly added maven to your environment variable.

To install and add maven to environment variable visit [Install maven](https://www.mkyong.com/maven/how-to-install-maven-in-windows/) if you are under a proxy check out [add proxy to maven](https://stackoverflow.com/a/47175285/8149916)

Navigate to the root of the project via command line and execute the command

```java
mvn spring-boot:run
```

The CLI will run your application on the configured port and you can access it just like you would if you start the app in an IDE.

Note: This will work only if you have maven added to your pom.xml

## common use command

```bash
mvn clean
mvn install
mvn package
mvn deploy
mvn dependency:tree
mvn dependency:analyze
mvn test
mvn compile
```