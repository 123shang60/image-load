# image-load

## 功能

在隔离环境k8s交付场景（无外网连接），通过s3存储动态为k8s节点load镜像

## 架构

server 端：

1. 为 agent 提供注册中心
2. 提交下载导入请求

agent 端：

- 基于 `docker client` 实现 docker 镜像的自动导入

## 调用实例

基于 `image-load-agent.yml` `image-load-server.yml` 进行部署

预先将要加载的 docker 镜像导入到 minio 中

使用如下命令：

```bash
curl --location --request POST 'localhost:8080/load' \
--header 'Content-Type: application/json' \
--data-raw '{
    "access_key": "admin",
    "secret_key": "123456789",
    "end_point": "192.168.31.10:9000",
    "bucket": "test",
    "item": "alpine.tar"
}'
```

一段时间后执行完毕，得到结果：

```json
[
    {
        "Name": "agent-local",
        "Code": 200,
        "Data": "ok!"
    }
]
```

## agent 列表获取

```bash
curl --location --request GET '127.0.0.1:8080/nodelist'
```

结果示例：

```json
{
    "agent-local": {
        "Object": {
            "name": "agent-local",
            "addr": "127.0.0.1",
            "port": "8081"
        },
        "Expiration": 1632124614618188000
    }
}
```

## 已知问题

1. 要求全部 k8s 节点的 docker api 版本必须完全一致，否则 agent 无法工作

## 改进方向

1. 服务内部通信应该基于 websocket 等长连接进行
2. 代码优化，部分重复代码重构
