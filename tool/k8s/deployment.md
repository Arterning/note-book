在生产环境中，我们基本上不会直接管理 pod，我们需要 `kubernetes` 来帮助我们来完成一些自动化操作，例如自动扩容或者自动升级版本。可以想象在生产环境中，我们手动部署了 10 个 `hellok8s:v1` 的 pod，这个时候我们需要升级成 `hellok8s:v2` 版本，我们难道需要一个一个的将 `hellok8s:v1` 的 pod 手动升级吗？

这个时候就需要我们来看 `kubernetes` 的另外一个资源 `deployment`，来帮助我们管理 pod

创建一个 `deployment.yaml` 的文件。来管理 `hellok8s` pod

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hellok8s-deployment
spec:
  replicas: 3 #pod 副本数量
  selector:
    matchLabels:
      app: hellok8s
  template:
    metadata:
      labels:
        app: hellok8s
    spec:
      containers:
        - image: guangzhengli/hellok8s:v1
          name: hellok8s-container
```

`kind` 表示我们要创建的资源是 `deployment` 类型， `metadata.name` 表示要创建的 deployment 的名字，这个名字需要是**唯一**的

```yaml
kubectl apply -f deployment.yaml

kubectl get deployments
#NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
#hellok8s-deployment   1/1     1            1           39s

kubectl get pods             
#NAME                                   READY   STATUS    RESTARTS   AGE
#hellok8s-deployment-77bffb88c5-qlxss   1/1     Running   0          119s

kubectl delete pod hellok8s-deployment-77bffb88c5-qlxss 
#pod "hellok8s-deployment-77bffb88c5-qlxss" deleted

kubectl get pods                                       
#NAME                                   READY   STATUS    RESTARTS   AGE
#hellok8s-deployment-77bffb88c5-xp8f7   1/1     Running   0          18s
```

如果我们的镜像升级到V2 修改配置文件即可

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hellok8s-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: hellok8s
  template:
    metadata:
      labels:
        app: hellok8s
    spec:
      containers:
        - image: guangzhengli/hellok8s:v2
          name: hellok8s-container
```

Rolling Update(滚动更新)

如果我们在生产环境上，管理着多个副本的 `hellok8s:v1` 版本的 pod，我们需要更新到 `v2` 的版本，像上面那样的部署方式是可以的，但是也会带来一个问题，就是所有的副本在同一时间更新，这会导致我们 `hellok8s` 服务在短时间内是不可用的，因为所有 pod 都在升级到 `v2` 版本的过程中，需要等待某个 pod 升级完成后才能提供服务。

这个时候我们就需要滚动更新 (rolling update)，在保证新版本 `v2` 的 pod 还没有 `ready` 之前，先不删除 `v1` 版本的 pod。

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hellok8s-deployment
spec:
  strategy:
     rollingUpdate:
      maxSurge: 1 #最大峰值，用来指定可以创建的超出期望 Pod 个数的 Pod 数量
      maxUnavailable: 1 #最大不可用，用来指定更新过程中不可用的 Pod 的个数上限
  replicas: 3
  selector:
    matchLabels:
      app: hellok8s
  template:
    metadata:
      labels:
        app: hellok8s
    spec:
      containers:
      - image: guangzhengli/hellok8s:v2
        name: hellok8s-container
```

`spec.strategy.type` 有两种选择:

- **RollingUpdate:** 逐渐增加新版本的 pod，逐渐减少旧版本的 pod。
- **Recreate:** 在新版本的 pod 增加前，先将所有旧版本 pod 删除。

大多数情况下我们会采用滚动更新 (RollingUpdate) 的方式