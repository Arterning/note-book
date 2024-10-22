在默认情况下，pnpm 会在全局保存每个包的单个副本，并通过符号链接来共享它们。这意味着，如果多个项目使用相同版本的 npm 包，pnpm 会尝试通过符号链接将这些包共享在一个位置，从而节省磁盘空间并提高效率。

这种方式与传统的 npm 和 yarn 不同，它们通常会在每个项目中分别保存每个包的副本，可能会占用更多的磁盘空间。

尽管 pnpm 的共享机制可以减少磁盘使用量，但需要注意以下几点：

1. **符号链接的限制：** 由于操作系统和文件系统的限制，符号链接不是在所有情况下都能正常工作。在某些情况下，特别是涉及一些构建工具或特殊的文件操作的软件包时，可能会遇到问题。

2. **潜在的兼容性问题：** 由于 pnpm 的共享机制与传统的包管理方式不同，一些依赖可能会有兼容性问题。尽管 pnpm 在处理大多数情况下都能正常工作，但偶尔可能会遇到一些特殊情况。

3. **工程设置：** pnpm 需要每个项目的配置文件来决定如何处理依赖。在某些情况下，这可能需要一些额外的配置工作。

总体而言，pnpm 的目标是在提供一种更加高效的依赖管理方式的同时，尽可能地减少磁盘使用。然而，使用 pnpm 时仍然需要谨慎，特别是在处理复杂的项目结构和依赖关系时，建议在小范围内进行测试以确保一切正常。

## npm 包安装失败“npm ERR！错误：EACCES：权限被拒绝

```bash
 pnpm config set unsafe-perm true
```

- https://github.com/microsoft/WSL/issues/14