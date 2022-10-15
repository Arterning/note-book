本文主要介绍在 MacOS 上使用 qemu 搭建 Linux Kernel 的开发环境。（在开始之前需要注意的是，本文中的 Linux 开发环境是一个远程服务器，而 qemu 被安装在本地的 MacOS 上。通常并不需要这样折腾，直接将 qemu 安装在 Linux 中更加方便，而且 qemu 是可以 -nographic 无图形界面运行的。）

1. 为什么需要 qemu?
qemu 是一个硬件虚拟化程序( hypervisor that performs hardware virtualization)，与传统的 VMware / VirtualBox 之类的虚拟机不同，它可以通过 binary translation 模拟各种硬件平台（比如在 x86 机器上模拟 ARM 处理器）。而 VirtualBox 等更多是通过虚拟化来进行资源隔离，以便在其上运行多个 guest os。
基于 qemu 的硬件模拟能力，我们可以轻松搭建指定硬件平台的运行实验环境。
qemu 与 VirtualBox 另一个不同点在于，在 VirtualBox 上必须安装一个完整的操作系统套件，而通过 qemu 我们可以通过参数直接启动到一个裸的 Linux Kernel，连 bootloader 都不需要关心。在此之外，按需配置相关工具套件与启动好的 Kernel 一起工作即可。
qemu 提供的这种高度可定制化的『白盒』能力，使得我们可以按需构建快速、轻量级的开发环境，提供流畅的开发体验。


构建一个压缩过的内核镜像
make bzImage

接下来，编译在配置阶段选择的内核模块：
make modules
编译好的内核模块 *.ko 文件存在于模块对应的源码目录中。


编译好内核以后，我们就可以使用 qemu 启动内核了。只需要使用 -kernel 参数告诉 qemu 内核文件的位置即可：
qemu-system-x86_64 \
    -m 512M \  # 指定内存大小
    -smp 4\  # 指定虚拟的 CPU 数量
    -kernel ./bzImage  # 指定内核文件路径

不出意外的话，就可以在启动窗口中看到内核的启动日志了。在内核启动的最后，会出现一条 panic 日志：
Kernel panic - not syncing: VFS: Unable to mount root fs on unknown-block(0, 0)
从日志内容可以看出，内核启动到一定阶段后尝试加载根文件系统，但我们没有指定任何磁盘设备，所以无法挂载根文件系统。而且上一节中编译出来的内核模块现在也没有用上，内核模块也需要存放到文件系统中供内核需要的时候进行加载。



创建磁盘镜像文件
使用 qemu-img 创建一个 512M 的磁盘镜像文件：
qemu-img create -f raw disk.raw 512M

现在 disk.raw 文件就相当于一块磁盘，为了在里面存储文件，需要先进行格式化，创建文件系统。比如在 Linux 系统中使用 ext4 文件系统进行格式化：
mkfs -t ext4 ./disk.raw

完整文章
//TODO qemu
https://www.jianshu.com/p/e1a4b5b808e0