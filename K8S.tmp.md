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

$ kubectl config use-context docker-desktop



$ goctl kube deploy -name nodejs -namespace keepfe -image ccr.ccs.tencentyun.com/keep/nodejs-build -o nodejs.yaml -port 6379
$ kubectl create namespace keepfe
$ kubectl apply -f nodejs.yaml
$ kubectl get all -n nodejs
$ kubectl run -i --tty --rm cli --image=nodejs -n keepfe -- sh
$ kubectl get pods -n keepfe
$ kubectl get pods -n keepfe -o yaml 
$ kubectl exec -it cli -n keepfe /bin/sh  #执行pod里面的 容器
$ kubectl exec -ti <your-pod-name>  -n <your-namespace>  -- /bin/sh
$ kubectl delete pod nodejs-577ccc9697-869j4 -n keepfe
$ kubectl get deployment -n keepfe
$ kubectl delete deployment nodejs -n keepfe


# 查看日志
$ kubectl -n default logs -f goweb
```

minikube
[参照链接](https://segmentfault.com/a/1190000018607114)
```shell
$ minikube ssh
  #docker ps | awk '{print $2F}'
$ kubectl config view
$ kubectl get node -o wide
  # 相当于用yaml 创建 pod
$ kubectl run goweb --image=hello  --port=8888 --replicas=3
$ kubectl create deployment goweb --image=192.168.3.85:5000/netkiller/welcome:latest
$ kubectl expose deployment goweb --port=8888 --target-port=8888 --type=NodePort
$ minikube service welcome --url
$ curl http://192.168.64.5:30257/from/you

# 报错指南
进入node节点去看
```

集群 节点 pod 服务



# 宿主本地源

``` json
{
  "insecure-registries" : [
    "10.21.71.237:5000"
  ],
  "debug" : true
}
```

```shell
$ docker pull registry
$ docker run -d -p 5000:5000 -v $(pwd):/var/lib/registry --restart always --name registry registry:2
$ docker tag hello:v1 10.2.17.5:5000/hello-local

```

```
minikube delete && minikube start --cpus=2 --memory=4096 --disk-size=10g \
  --driver=hyperkit \
  # --docker-env http_proxy=10.2.17.5:1087 \
  # --docker-env https_proxy=10.2.17.5:1087 \
  # --docker-env no_proxy=192.168.99.0/24 \
  --registry-mirror=https://registry.docker-cn.com \
  --insecure-registry=10.2.17.5:5000 \
  --service-cluster-ip-range='10.10.0.0/24'
```

#https://zhuanlan.zhihu.com/p/261722859