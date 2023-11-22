您可以使用Spring Boot提供的命令行工具Spring Initializr（或`spring init`命令）快速创建一个Spring Boot项目。以下是通过命令行工具创建Spring Boot项目的基本步骤：

1. **确保安装Spring Boot CLI**：

   首先，确保您已经安装了Spring Boot CLI。您可以从Spring官方网站（https://spring.io/projects/spring-boot#learn）
   下载并安装它，根据您的操作系统选择适当的安装方法。

2. **使用Spring Initializr创建项目**：

   在命令行中，您可以使用`spring init`命令来创建Spring Boot项目。以下是一个基本的示例命令：

   ```bash
   spring init --dependencies=web,data-jpa,thymeleaf my-spring-boot-project
   ```

   这个命令创建一个名为`my-spring-boot-project`的新Spring Boot项目，并添加了一些常用的依赖，包括`web`（用于构建Web应用程序）、`data-jpa`（用于JPA数据访问）、`thymeleaf`（用于模板引擎）。您可以根据您的需求自定义项目依赖。

3. **进入项目目录**：

   使用`cd`命令进入新创建的项目目录：

   ```bash
   cd my-spring-boot-project
   ```

4. **运行项目**：

   使用以下命令启动Spring Boot应用程序：

   ```bash
   ./mvnw spring-boot:run
   ```

   或者如果您在Windows上使用CMD或PowerShell，则可以运行：

   ```bash
   mvnw spring-boot:run
   ```

   Spring Boot应用程序将启动并运行在默认端口（通常是8080）上。

现在，您已经成功创建了一个新的Spring Boot项目，并可以根据您的需求在其中编写代码。可以根据项目需要添加更多的依赖、配置、控制器、服务等。要在IDE中打开项目进行开发，您可以使用您喜欢的集成开发环境（如Eclipse、IntelliJ IDEA等）导入项目。


## 指定版本
Spring Boot CLI 可以使用 `-v` 或 `--version` 选项来指定要创建的 Spring Boot 项目的版本。您可以通过以下方式在创建项目时指定版本：

```shell
spring init --dependencies=web,data-jpa,thymeleaf -v <版本号> my-spring-boot-project
```

或者使用 `--version`：

```shell
spring init --dependencies=web,data-jpa,thymeleaf --version <版本号> my-spring-boot-project
```

在上述命令中，`<版本号>` 应该替换为您想要使用的 Spring Boot 版本号。例如，要创建一个使用 Spring Boot 2.5.4 版本的项目，可以运行：

```shell
spring init --dependencies=web,data-jpa,thymeleaf -v 2.5.4 my-spring-boot-project
```

这将创建一个新的 Spring Boot 项目，并使用指定版本的 Spring Boot。确保指定的版本号是有效的，并且与您计划使用的依赖项兼容。
