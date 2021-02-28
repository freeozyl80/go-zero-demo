# 这里想想分享的内容

### stage1

1. goctl 声明API

```shell
    goctl api -o shorturl.api
```

.api 文件

```code
@server // server属性
handler 定义服务名称
get/post/ 等方法声明

```

2. goctl 生成 API Gateway 代码

``` shell
goctl api go -api shorturl.api -dir .

```

- 运行
``` shell
curl -i "http://localhost:8888/shorten?url=http://www.xiaoheiban.cn"
```
> 备注
  shorturl-api.yaml 的配置项


3. 大概理一下代码：
 - main.go
    * svc 创建context
    * ret 创建server
 - handler > router.go
    * 给server注册路由handler
 - handler > ${name}handler.go
    * logic context
    * logic req && res (from types)

### stage2

4. rpc 目录创建
```shell
goctl rpc template -o transform.proto
```

```
goctl rpc proto -src transform.proto -dir .
```
- 运行
```shell
etcdctl get transform.rpc --prefix
```

5. api实现 rpc 调用的主要思路

    * 配置文件更新 yaml
    * config 添加 Transform 服务依赖
    * svc 里面的 context 修改 //l.svcCtx.Transformer.Expand
    * logic 逻辑更新
    

### stage3

6. 增加数据库的调用

```shell
# -c 表示使用redis cache
goctl model mysql ddl -c -src shorturl.sql -dir .
```

7. config 配置里面 指定 数据库 

> api 里面的配置增加的是 rpc调用
> rpc 配置里面增加的是 数据库调用

8. rpc 接入数据库支持整个流程
    * 配置文件更新 yaml
    * config 添加 数据库 服务依赖
    * svc 里面的 context 修改   // l.svcCtx.Model.Insert
    * logic 逻辑更新

### stage4 补充一个生成ts 的 request 的代码

```

```