> nodejs 报错 Client does not support authentication protocol requested by server; consider upgrading MySQL client，mysql 版本是 8 如何解决？

---

该错误通常出现在 MySQL 8 中使用旧版 MySQL 客户端的情况下，因为 MySQL 8 更改了默认身份验证插件。解决此问题的一种方法是将 MySQL 客户端升级到支持新的身份验证插件的版本。

以下是一些可能的解决方法：

1.  升级 MySQL 客户端：

确保使用支持新的身份验证插件的 MySQL 客户端版本，如 MySQL 8.0.4 或更高版本。

2.  更改 MySQL 用户的身份验证插件：

如果您无法升级 MySQL 客户端，则可以更改 MySQL 用户的身份验证插件。例如，可以使用以下命令更改现有用户的身份验证插件：

sql

```sql
ALTER USER 'user'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
```

3.  更改 MySQL 服务器的默认身份验证插件：

如果您无法升级 MySQL 客户端或更改用户的身份验证插件，则可以尝试更改 MySQL 服务器的默认身份验证插件。例如，可以在 MySQL 配置文件中添加以下行：

javascript

```javascript
default-authentication-plugin=mysql_native_password
```

这将使 MySQL 服务器使用 mysql_native_password 身份验证插件作为默认插件。

请注意，更改 MySQL 用户或服务器的身份验证插件可能会对安全性产生影响。因此，在更改身份验证插件之前，请确保了解可能的风险和限制，并采取必要的安全措施。

---

> nodejs 如何升级 MySQL 客户端？需要升级 mysql 包到什么版本？

---

要升级 Node.js 中使用的 MySQL 客户端，您需要升级 `mysql` 包到支持新的 MySQL 8 身份验证插件的版本。具体来说，您需要升级 `mysql` 包至版本 2.0.0 或更高版本，因为从此版本开始，`mysql` 包已经支持 MySQL 8 的新身份验证插件。

要升级 `mysql` 包，您可以使用 npm 命令：

java

```java
npm install mysql@latest
```

这将安装最新版本的 `mysql` 包。如果您希望安装特定版本的 `mysql` 包，则可以使用类似以下命令：

css

```css
npm install mysql@2.18.1
```

这将安装版本号为 2.18.1 的 `mysql` 包。

请注意，如果您在 Node.js 14 或更高版本中使用 `mysql` 包，则可能会遇到一些问题，因为从 Node.js 14 开始，`Buffer` 构造函数已被弃用。如果您遇到此问题，则可以升级到 `mysql2` 包，它是一个兼容 `mysql` 包的替代品，并支持 Node.js 14 和更高版本。您可以使用以下命令安装 `mysql2` 包：

`npm install mysql2`

然后，您可以将 `mysql` 包替换为 `mysql2` 包，并更新您的代码以使用 `mysql2` 包提供的 API。

---

> npm 中的--legacy-peer-deps 是什么意思？

---

`--legacy-peer-deps` 是 npm 7 中的一个选项，它用于在安装依赖项时保留旧版本的依赖项解析行为。在默认情况下，npm 7 会使用更严格的对等依赖项解析规则，这可能会导致某些软件包无法安装或编译失败。

当您在使用 npm 7 安装软件包时，如果某些软件包依赖于旧版本的软件包，但是这些软件包不符合新的对等依赖项解析规则，npm 7 可能会拒绝安装这些软件包并显示错误消息。在这种情况下，您可以使用 `--legacy-peer-deps` 选项，以便 npm 7 在安装依赖项时使用旧版本的依赖项解析行为。

使用 `--legacy-peer-deps` 选项时，npm 7 将使用与 npm 6 相同的对等依赖项解析规则，这将允许旧版本的软件包与新版本的软件包一起使用，从而解决可能出现的依赖项解析问题。请注意，使用 `--legacy-peer-deps` 选项可能会导致某些软件包存在安全漏洞或依赖项冲突问题。因此，您应该在使用此选项时格外注意安全性和稳定性问题，并尽可能升级软件包以使用新的对等依赖项解析规则。

---

> 如果旧版本的软件包与新版本的软件包一起使用，程序会使用旧版本的软件包还是新版本的软件包？

---

如果您在使用 `--legacy-peer-deps` 选项时安装软件包，而这些软件包依赖于旧版本和新版本的软件包，程序将使用旧版本的软件包，而不是新版本的软件包。

`--legacy-peer-deps` 选项的作用是使 npm 7 在解析依赖关系时使用与 npm 6 相同的规则，这样旧版本的软件包和新版本的软件包都可以与其他软件包一起工作。这意味着，如果某个软件包依赖于旧版本的软件包，而您使用 `--legacy-peer-deps` 安装该软件包，那么 npm 7 将使用旧版本的软件包，而不是新版本的软件包。

请注意，尽管 `--legacy-peer-deps` 选项可以解决某些依赖项问题，但它并不总是最好的解决方案。使用 `--legacy-peer-deps` 选项可能导致您的应用程序存在安全漏洞或其他问题，因此您应该谨慎使用，并尽可能升级软件包以使用新版本的依赖项解析规则。

---

> .nvmrc 文件的作用？

---

`.nvmrc` 文件是用于 Node Version Manager（NVM）的配置文件，用于指定应该使用哪个 Node.js 版本来运行特定的项目或应用程序。

当您在使用 NVM 管理多个 Node.js 版本时，您可能需要在不同的项目或应用程序之间切换版本。为了方便起见，您可以在每个项目的根目录中创建一个 `.nvmrc` 文件，并在其中指定应该使用的 Node.js 版本号。

例如，如果您在一个项目中使用 Node.js 14.x 版本，您可以在项目的根目录中创建一个名为 `.nvmrc` 的文件，并在其中写入以下内容：

`14`

这将告诉 NVM 在运行该项目之前使用 Node.js 14.x 版本。然后，您可以在该项目的根目录中运行 `nvm use` 命令，NVM 将读取 `.nvmrc` 文件并自动切换到指定的 Node.js 版本。

`.nvmrc` 文件的优点是它能够方便地让您在不同的项目或应用程序之间切换 Node.js 版本，而无需手动更改全局 Node.js 版本或在项目中手动指定版本号。这可以节省您的时间和精力，并确保每个项目都使用适当的 Node.js 版本。

---
