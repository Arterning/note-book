/**
 * 
 * 
 * cargo下载依赖的命令是哪个？
 * 
 * 在 Rust 的项目中，使用 cargo 下载依赖的命令是：

cargo build

这个命令会读取 Cargo.toml 文件中的依赖关系，并从 crates.io 下载依赖的库，最后编译整个项目。



下载的依赖包的存放位置是在哪里?

下载的依赖包存放在 $HOME/.cargo/registry 目录下面，这个目录是用户级别的，因此可以被所有项目共享。

你也可以在某个项目目录下面看到本地编译生成的文件，它们存放在 target 目录下面。

需要注意的是，如果你使用了不同的 Rust 版本，那么你可能会在 $HOME/.cargo 目录下面看到不同的文件夹，例如：$HOME/.cargo/registry/src 和 $HOME/.cargo/registry/{hash}/src。这是因为 cargo 会根据 Rust 的版本为不同的项目生成不同的依赖关系。

 */