# image-load

## 功能

在隔离环境k8s交付场景（无外网连接），通过s3存储动态为k8s节点load镜像

## 架构

server 端：

1. 为 agent 提供注册中心
2. 提交下载导入请求

agent 端：

- 基于 `docker client` 实现 docker 镜像的自动导入