# maven-wrapper.jar 和 maven-wrapper.properties 作用是什么？

maven-wrapper.jar 和 maven-wrapper.properties 是 Maven Wrapper 的两个核心文件，它们是用于简化 Maven 的版本管理的工具。

maven-wrapper.jar 是一个 Java 程序，它是用来管理 Maven 的版本的，因此保证了项目中使用的是特定版本的 Maven。
maven-wrapper.properties 文件存储了 Maven Wrapper 应该使用的 Maven 版本信息。它包含了 Maven 版本的 URL 以及需要的密钥。
使用 Maven Wrapper 可以确保项目中所使用的 Maven 版本与预期一致，并且能够避免因为不同的 Maven 版本导致的问题。团队成员可以确保在开发项目时使用相同版本的 Maven，因此能够保证项目在不同的开发环境中的一致性。

# 传统使用 maven 的方式

先到官网上下载 maven,然后配置环境变量把 mvn 可执行文件路径加入到环境变量，以便之后使用直接使用 mvn 命令。
另外项目 pom.xml 文件描述的依赖文件默认是下载在用户目录下的.m2 文件下的 repository 目录下。
如果需要更换 maven 的版本，需要重新下载 maven 并替换环境变量 path 中的 maven 路径。

# maven-wrapper 的目的

maven-wrapper 的出现就是为了在更换 maven 版本时不用手动去做上面的操作:

执行 mvnw 比如 mvnw clean，如果本地没有匹配的 maven 版本，直接会去下载 maven，放在用户目录下的.m2/wrapper 中
并且项目的依赖的 jar 包会直接放在项目目录下的 repository 目录，这样可以很清晰看到当前项目的依赖文件。
如果需要更换 maven 的版本，只需要更改项目当前目录下.mvn/wrapper/maven-wrapper.properties 的 distributionUrl 属性值，更换对应版本的 maven 下载地址。mvnw 命令就会自动重新下载 maven。

可以说带有 mvnw 文件的项目，除了额外需要配置 java 环境外，只需要使用本项目的 mvnw 脚本就可以完成编译，打包，发布等一系列操作。

# 个人理解

maven wrapper 可以自动下载 maven，但实际上我们常用的 idea 软件都自带了 maven。
如果用上了 ide，一般习惯也是直接使用 Navigation Bar 执行 maven 命令比较方便。
maven wrapper 根据配置自动切换 maven 版本。这个看起来很有用，但实际上 maven 版本也是很稳定。很少会出现需要切换 maven 版本的情况
使用 mvnw 命令会在直接当前项目下生成 repository，看起来每一个项目独立了 repository，很模块化的样子。但是这样不仅浪费了磁盘空间，且实际上开发中并不关心 repository，idea 会自动有 external librayies 目录提供查看依赖的 jar 包。

maven-wrapper 解决了 2 个问题

可以为某个 Java 工程指定某个特定 Maven 版本，避免因为版本差异引起的诡异错误，这样就统一该项目的开发环境
不再需要提前安装 Maven，简化了开发环境的配置

# More

mvnw.cmd 是 Maven Wrapper 的一个命令行工具，它是在 Windows 系统上使用的。该命令的作用是简化 Maven 的版本管理，并确保项目中所使用的 Maven 版本是特定的。

通过使用 mvnw.cmd，开发人员可以在 Windows 系统上执行 Maven 的常用任务，如编译、测试、打包等，而不必担心在不同的机器上使用的 Maven 版本不一致的问题。该工具会使用预定义的 Maven 版本，从而确保项目的一致性和可重复性。

例如，要在 Windows 系统上编译项目，只需要在命令行中执行以下命令：

Copy code
mvnw.cmd clean install
这个命令会在 Windows 系统上自动下载并使用预定义的 Maven 版本，并使用该版本执行 clean 和 install 任务。
