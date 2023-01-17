/**
 * 
 * 市面上常见的日志框架有很多，它们可以被分为两类：日志门面（日志抽象层）
 * 为 Java 日志访问提供一套标准和规范的 API 框架，其主要意义在于提供接口。
 * 
 * 和日志实现 Log4j、JUL（java.util.logging）、Log4j2、Logback
Spring Boot 选用 SLF4J + Logback 的组合来搭建日志系统。
 */

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Log {
    


    testlog() {
        Logger logger = LoggerFactory.getLogger(HelloWorld.class);
        //调用 sl4j 的 info() 方法，而非调用 logback 的方法
         logger.info("Hello World");
    }
}
