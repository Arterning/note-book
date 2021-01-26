1. install tool..
sudo apt-get install git fakeroot build-essential ncurses-dev xz-utils libssl-dev bc flex libelf- dev bison

2. 从当前机器的启动⽬录拷⻉配置信息到源代码⽬录。这步操作的意思是我们编译内核的配置采⽤ ⽤当前环境⼀致的配置。
cp -v /boot/config-$(uname -r) .config

3. start menu config tool
make menuconfig

4. start compile
make -j 10 上⾯参数是并发数量，通常可以是CPU的2倍。

5. make && sudo make install

6. 查看内核版本
uname -mrs


7. error
  CALL    scripts/checksyscalls.sh
  CALL    scripts/atomic/check-atomics.sh
  DESCEND objtool
  CHK     include/generated/compile.h
  CHK     kernel/kheaders_data.tar.xz
make[1]: *** No rule to make target 'debian/canonical-certs.pem', needed by 'certs/x509_certificate_list'.  Stop.
make: *** [Makefile:1858: certs] Error 2

vim .config
然后找到
CONFIG_SYSTEM_TRUSTED_KEYS,将其设置为空,也就是下面这个样子。
CONFIG_SYSTEM_TRUSTED_KEYS=””