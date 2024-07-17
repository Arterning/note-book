
ingress也是k8s api的标准资源类型之一，它其实就是一组基于DNS名称（host）或URL路径把请求转发到指定的service资源的规则。用于将集群外部的请求流量转发到集群内部完成的服务发布。我们需要明白的是，ingress资源自身不能进行“流量穿透”，仅仅是一组规则的集合，这些集合规则还需要其他功能的辅助，比如监听某套接字，然后根据这些规则的匹配进行路由转发，这些能够为ingress资源监听套接字并将流量转发的组件就是ingress controller。



ingress控制器不同于deployment等pod控制器的是，ingress控制器不直接运行为kube-controller-manager的一部分，它仅仅是k8s集群的一个附件，类似于coreDNS,需要在集群上单独部署。

ingress controller通过监视api server获取相关ingress、service、endpoint、secret、node、configmap对象，并在程序内部不断循环监视相关service是否有新的endpoint变化，一旦发生变化则自动更新nginx.conf模板配置并产生新的配置文件进行reload。




总结：

ingress 就是根据我们配置的yaml生成的规则，是一组资源的映射规则 

```

域名，路径 ==> service 资源
```


ingress controller 监听ingress service的变化，一旦发生变化更新nginx.conf生成新的配置文件并且reload



  
