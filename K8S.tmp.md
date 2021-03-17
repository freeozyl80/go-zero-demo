### kubelet

kubelete 作为节点代理，每个节点都有kubelete金川，处理master下发的任务（Podspec:yaml => pod）

容器检测 && 容器健康管理

### kukbeadm 

用来部署 kubernetes 集群
组件间访问授权，健康检查


### 备注：

Pod 
可以容纳多个容器
相同的pod中的任何容器共享相同名称空间和本地网络

pod 由 deployment（部署） 来进行管理
：
声明一个pd应该同时运行多个副本
不必手动处理pod，只需声明系统的期望状态

### 尝试使用kubernete 对象

``` shell
$ goctl kube deploy -name nodejs -namespace keepfe -image ccr.ccs.tencentyun.com/keep/nodejs-build -o nodejs.yaml -port 6379
$ kubectl create namespace keepfe
$ kubectl apply -f nodejs.yaml
$ kubectl get all -n nodejs
$ kubectl run -i --tty --rm cli --image=nodejs -n keepfe -- sh
$ kubectl get pods -n keepfe
$ kubectl get pods -n keepfe -o yaml 
$ kubectl exec -it cli -n keepfe /bin/sh  #执行pod里面的 容器
$ kubectl delete pod nodejs-577ccc9697-869j4 -n keepfe
$ kubectl get deployment -n keepfe
$ kubectl delete deployment nodejs -n keepfe

```

minikube
[参照链接](https://segmentfault.com/a/1190000018607114)
```shell
$ minikube ssh
  #docker ps | awk '{print $2F}'
$ kubectl config view
$ kubectl get node -o wide
```

集群 节点 pod 服务


