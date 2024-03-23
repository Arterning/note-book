## minikube 命令


```
# 启动集群
minikube start
# 查看节点。kubectl 是一个用来跟 K8S 集群进行交互的命令行工具
kubectl get node
# 停止集群
minikube stop
# 清空集群
minikube delete --all
# 安装集群可视化 Web UI 控制台
minikube dashboard
```

https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/nd7yOvdY




## 检查自己的镜像是否有问题

```
docker run -p 3000:3000 --name test-app --env MONGO_URL="mongodb://192.168.1.128:27017/nest-domain-resolver" arterning/nest-domain-resolver:v1
```


## 下载helm

每一个版本的 Helm 提供多种操作系统的二进制版本，支持选择 Linux、Mac OS、Windows 平台，以及 amd64、arm、i386、ppc64le、s390x 的 CPU 类型。

这些二进制版本可以手动下载和安装。

下载地址：

- [https://github.com/helm/helm/releases](https://github.com/helm/helm/releases)




## helm 安装本地chart

在解压缩后的文件夹中，您将找到 Chart.yaml 和 values.yaml 文件。使用以下命令安装 Chart：

```yaml
helm install release_name ./chart_folder --values ./chart_folder/values.yaml
```

其中，release_name 是您为安装创建的名称，./chart_folder 是您解压缩 Chart 文件的文件夹路径。如果 Chart 中有任何自定义值，可以使用 --values 标志并指定 values.yaml 文件的路径。



验证 Chart 是否已安装：

可以使用以下命令验证 Chart 是否已成功安装：

```
helm ls

```


```
helm install domain-resolver . --values .\values.yaml

helm uninstall domain-resolver
```