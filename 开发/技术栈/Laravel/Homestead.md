

重新构建以生效

If you change the sites property after provisioning the Homestead virtual machine, you should execute the vagrant reload --provision command in your terminal to update the Nginx configuration on the virtual machine.


Homestead scripts are built to be as idempotent as possible. However, if you are experiencing issues while provisioning you should destroy and rebuild the machine by executing the vagrant destroy && vagrant up command.


如果在配置 Homestead 虚拟机后更改站点属性，则应在终端中执行 vagrant reload --provision 命令来更新虚拟机上的 Nginx 配置。 Homestead 脚本被构建为尽可能幂等。


但是，如果您在配置时遇到问题，您应该通过执行 vagrant destroy && vagrant up 命令来销毁并重建机器。