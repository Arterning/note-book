

创建Pod和获取Pod

```
kubectl apply -f nginx.yaml        

kubectl get pods
# nginx-pod         1/1     Running   0           6s


# 将 `nginx` 默认的 `80` 端口映射到本机的 `4000` 端口
kubectl port-forward nginx-pod 4000:80

# 进入pod
kubectl exec -it nginx-pod /bin/bash


# delete pod

kubectl delete pod nginx-pod
# pod "nginx-pod" deleted

kubectl delete -f nginx.yaml
# pod "nginx-pod" deleted

```



创建和获取Deployment

```
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


Pod是Kubernetes集群进行管理的最小单元，它是容器运行的载体。 Pod可以认为是容器的封装，一个Pod中可以存在一个或多个容器。 这些容器共享存储、网络和运行资源。 由于Pod具有这些共享资源，因此它们可以作为一个整体进行调度和部署。




Deployment 提供的特性

 多副本，存活探针，就绪探针，滚动升级


Service 提供的特性

故障转移， 负载均衡，提供稳定的访问地址




## 注意

注意，无论你配置的service类型是clusterIP还是nodePort

你需要访问的都是minikube ip。。




## NodePort的工作原理

NodePort的工作原理其实就是**将service的端口映射到Node的一个端口上**，然后就可以通过`NodeIp:NodePort`来访问service了。



```
不知道什么原因，我本地的minikube ip是访问不到的
```