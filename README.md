## 这里记一下 go-zero 的相关一些知识点


#### goctl

1. 定义api 请求
2. 根据定义的api自动生成golang, typescript 等各种语言的带啊嘛
3. 生成mysql
4. 生成mongodb


>
  定义API文件

  Request:
  Response:
  service:
    @handeler
    URI
  
 直接通过 api 文件 生成接口：
 goctl api go -api shorturl.api -dir .



goctl rpc template -o transform.proto

goctl rpc proto -src transform.proto -dir .


----- etcd

配置管理
服务注册发现
选主
应用调度
分布式队列
分布式锁



// 需要安装etcd, 这里插播一下etcd 的教程



goreman  类似于nodemon


etcd 的常用命令：

etcdctl member list  //查看本地集群的服务器列表
etcdctl put foo "Hello World"
etcdctl get foo // --prefix
etcdctl del foo 
etcdctl watch foo
etcdctl lock foo
etcdctl txn -i //从标准库读取多个请求，看作一个原子性的事务进行执行

```shell
 etcdctl put user frank #设置user为 frank
 etcdctl txn -i #开启事务（-i表示交互模式）
 value("user") = "frank" #此命令是比较 user 的值与 frank 是否相等
 put result ok  #成功的条件
 put result failed #失败的条件
 
 etcdctl get result

```


etcd 设置过期时间
etcdctl lease grant 30    #设置30s的过期时间
etcdctl put --lease=${上一个的值} foo bar #应用过期时间
etcdctl lease revoke ${上一个的值} #撤销过期时间

etcdctl lease keep-alive ${上一个的值}
etcdctl lease timetolive ${上一个的值} 查看剩余的时间  --keys 看key值


> 服务发现实战：

整个代码的逻辑很简单，worker 启动时向 etcd 注册自己的信息，并设置一个带 TTL 的租约，每隔一段时间更新这个 TTL，如果该 worker 挂掉了，这个 TTL 就会 expire 并删除相应的 key。发现服务监听 workers/ 这个 etcd directory，根据检测到的不同 action 来增加，更新，或删除 worker。



大概参考这篇文章吧：https://zhuanlan.zhihu.com/p/96428375?from_voters_page=true




