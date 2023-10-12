# deploy fabric  
  
## 安装k8s集群  
```bash  
  
# 为每个节点分别设置对应主机名 假设是一主四从  
# master node1 node2 node3 node4  
hostnamectl set-hostname master  
hostnamectl set-hostname node1  
hostnamectl set-hostname node2  
hostnamectl set-hostname node3  
hostnamectl set-hostname node4  
  
# 所有节点都修改 hostsvim /etc/hosts  
master_ip master  
node1_ip node1  
node2_ip node2  
node3_ip node3  
node4_ip node4  
  
  
# 在所有节点运行 安装 dockertar -xzvf deploy_env.tar.gz && cd deploy_env  
./env.sh installenv -t centos -p /mydocker  
  
# 在所有节点运行 安装基础依赖包  
tar -xzvf rpm_pack.tar.gz && cd rpm_pack  
./install.sh  
echo 1 > /proc/sys/net/bridge/bridge-nf-call-iptables  
modprobe dm_thin_pool  
  
# 在所有节点运行 安装k8s node  
tar -xzvf kubernetes_env_pack.tar.gz && cd kubernetes_env_pack  
./kubernetes.sh releaseKubernetesEnv  
  
# 在所有节点运行 启动 kubelet、docker，并设置开机启动  
systemctl enable kubelet  
systemctl start kubelet  
systemctl enable docker  
systemctl start docker  
  
# 用 kubeadm 初始化集群（仅在主节点跑）  
./kubernetes.sh initKubernetes -i master_ip -m notup -t flannel  
  
# 把工作节点加入集群（只在工作节点跑）  
kubeadm token create --print-join-command  
kubeadm join master_ip:6443 --token xxx --discovery-token-ca-cert-hash xxx  
  
# 查看 Kubernetes 集群  
kubectl get nodes  
```  
  
  
  
  
## GlusterFS和Heketi  
Heketi（也称为Heketi Rest Server）是一个开源的存储管理器，主要用于管理和自动化分布式存储系统中的卷（volumes）。它最初由Gluster社区开发，并用于管理GlusterFS分布式文件系统中的卷。然而，它不仅限于GlusterFS，还可以与其他分布式存储后端一起使用。  
  
Heketi的目标是简化分布式存储系统的管理，尤其是在需要动态调整存储容量和卷分配时。通过提供RESTful API和命令行界面，管理员可以轻松地创建、删除和调整存储卷，而无需直接与底层存储后端交互。  
  
Heketi通过与存储后端（如GlusterFS、Ceph等）通信，自动管理卷的分布和平衡，确保数据在存储集群中的均匀分布和高可用性。这对于大规模的分布式存储环境非常有用，因为手动管理这些任务可能会变得非常繁琐和复杂。  
  
总而言之，Heketi是一个用于管理分布式存储系统中存储卷的工具，旨在简化存储管理和卷分配的过程。  
  
  
GlusterFS是一个开源的分布式文件系统，旨在提供可扩展的网络存储解决方案。它允许将多台服务器（节点）的存储资源组合在一起，形成一个统一的、分布式的文件存储池。这使得数据可以在这些节点之间自动分布和复制，从而提供高可用性、高性能和可扩展性。  
  
GlusterFS的工作方式是通过将各个存储节点之间的本地文件系统卷合并成一个全局的命名空间。这意味着你可以像操作本地文件一样操作分布在整个存储集群中的文件。数据可以根据需要进行分布、复制和条带化，从而实现高性能和冗余存储。  
  
GlusterFS提供了多种不同的卷类型，可以根据特定的需求和用例进行配置。一些常见的卷类型包括分布式卷、复制卷、条带卷和分布式复制卷等。  
  
GlusterFS的优点包括：  
  
1. **可扩展性：** 可以根据需要添加更多的存储节点，从而扩展存储容量和性能。  
2. **高可用性：** 数据可以在多个节点之间复制，确保数据的冗余和高可用性。  
3. **容错性：** 在节点故障时，数据仍然可用，因为数据在多个节点上复制。  
4. **灵活性：** 可以根据应用程序的需求配置不同类型的卷，从而平衡性能和容错需求。  
5. **分布式：** 提供了统一的全局命名空间，使得数据在整个存储集群中分布。  
  
总的来说，GlusterFS是一个强大的分布式文件系统，适用于需要高性能、高可用性和可扩展性的存储解决方案。它被广泛用于数据存储、云计算环境和大规模数据分析等场景。  
  
## 清理heketi和glusterfs  
如果在此之前有部署过glusterfs或者需要重新部署，可以先完全删除之前的服务和数据  
  
### 在master上删除heketi、glusterfs服务  
```yml  
kubectl delete -f heketi-bootstrap.json  
kubectl delete -f glusterfs-daemonset.json  
```  
### 在node上清理glusterfs数据  
如果我们希望在/dev/sdb (100G)上保存glusterfs数据 需要先执行如下命令  
```bash  
sudo rm -rf /var/lib/heketi  /run/lvm/  /etc/glusterfs/  /var/lib/glusterdsudo dmsetup statussudo dmsetup remove_allsudo dd if=/dev/zero of=/dev/sdb bs=100G count=1sudo wipefs -af /dev/sdbsudo dd if=/dev/zero of=/dev/sdb bs=100G count=1sudo wipefs -af /dev/sdbsudo dd if=/dev/zero of=/dev/sdb bs=100G count=1```  
  
## 部署heketi + glusterFs服务  
```bash  
tar -xzvf heketi-gluster.tar.gz && cd heketi-gluster  
# 节点贴标签  
kubectl label node node1 storagenode=glusterfskubectl label node node2 storagenode=glusterfskubectl label node node3 storagenode=glusterfskubectl label node node4 storagenode=glusterfs  
# 查看节点标签是否设置成功  
kubectl get nodes --show-labels  
vi glusterfs-daemonset.json，修改imagePullPolicy:IfNotPresent  
kubectl apply -f glusterfs-daemonset.jsonkubectl create -f heketi-service-account.jsonkubectl create clusterrolebinding heketi-gluster-admin --clusterrole=edit --serviceaccount=default:heketi-service-accountkubectl create secret generic heketi-config-secret --from-file=./heketi.jsonvi heketi-bootstrap.json，修改 imagesPullPolicy:IfNotPresentkubectl apply -f heketi-bootstrap.jsonkubectl get pods  
  
kubectl exec -it deploy-heketi-xxxx-xxx bashexport HEKETI_CLI_SERVERE=http://heketi_ip:8080heketi-cli -s $HEKETI_CLI_SERVERE --user admin --secret 'My Secre' topology infovi topology.json，修改hostname IP  mountPath  
heketi-cli -s $HEKETI_CLI_SERVERE --user admin --secret 'My Secre' topology load --json=topology.jsonheketi-cli -s $HEKETI_CLI_SERVERE --user admin --secret 'My Secre' topology info  
heketi-cli -s $HEKETI_CLI_SERVERE --user admin --secret 'My Secre' setup-openshift-heketi-storage Saving heketi-storage.jsonctrl+D  
  
kubeclt cp deploy-heketi-xxx-xxx:/home/heketi-storage.json ./heketi-storage.jsonkubeclt apply -f heketi-storage.jsonkubectl delete all,svc,jobs,deployment,secret --selector=”deploy-heketi”  
kubectl apply -f heketi-deployment.jsonvi storageClass.yaml，修改IP  
kubectl apply -f storageClass.yaml  
  
# 查看k8s集群状态  
kubectl get all -A  
```  
  
## 部署crchain  
```yml  
tar -xzvf crchain_image.tar.gz && cd crchain_image  
./crchain_images.sh release all  
  
tar -xzvf crchain.tar.gz && cd crchain  
  
# -i baas平台nginx ip地址；-k k8s monitor地址(dashboard地址)； -a external地址  
./deploy.sh generate -s all -i nginx_ip -w 3001 -k dashboard_ip -e 8200 -a 10.200.207.11  
  
./deploy.sh up -s all  
  
```  
  
  
## 添加集群  
IP地址为k8s的nginx-ingress所在主机的IP  
  
集群文件为k8s master机器的/etc/kubernetes/admin.conf的内容


# 数据上链  
  
定义基本数据结构  
  
```go  
type FarmBatch struct {  
   FarmList []Farm `json:"farm_list"` // 手动批量农场信息  
}  
  
type Farm struct {  
   TenantId string `json:"tenant_id"` // 公司ID  
   FarmId   string `json:"farm_id"`   // 农场id  
   SectionIdList []string `json:"section_id_list"` // 地块列表  
   SoilTestReportList  []TestReport `json:"soil_test_report_list"`  // 检测报告  
   WaterTestReportList []TestReport `json:"water_test_report_list"` // 检测报告  
   Status              int          `json:"status"`                 // 数据状态，0正常，1删除。  
}  
  
```  
  
接受请求体，反序列化，上传到区块链存储  
  
```go  
func PutFarm(c *gin.Context) {  
   //farm := &define.Farm{}   farmBatch := &define.FarmBatch{}   body, err := ioutil.ReadAll(c.Request.Body)   if err != nil {      errStr := fmt.Sprintf("get body error %v", err)      utils.Response(errStr, c, http.StatusBadRequest)   }   if err := json.Unmarshal(body, farmBatch); err != nil {      errStr := fmt.Sprintf("json unmarshal error %v", err)      utils.Response(errStr, c, http.StatusBadRequest)   }   if len(farmBatch.FarmList) == 0 {      return   }   txId := ""   for _, farm := range farmBatch.FarmList {      data, err := json.Marshal(farm)      if err != nil {         errStr := fmt.Sprintf("json marshal error %v", err)         utils.Response(errStr, c, http.StatusBadRequest)      }      // 更新链上数据  
      info := []string{"PutFarm", string(data)}      resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")      if err != nil {         errStr := fmt.Sprintf("数据上链失败 %v", err)         utils.Response(errStr, c, http.StatusBadRequest)      }      txId = resp.TxID   }   utils.Response(txId, c, http.StatusOK)   return}  
```