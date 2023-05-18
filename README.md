# chat_azure --- Aure Open AI API Transparent Proxy

[English README](./README_en.md)

这是一个带有用户管理功能的 Azure Open AI API 透明代理后台。此项目基于以下库：

- [libli/chat](https://github.com/libli/chat)
- [haibbo/cf-openai-azure-proxy](https://github.com/haibbo/cf-openai-azure-proxy)

## 特性

- 用户身份验证和管理
- Azure Open AI API 的透明代理

## 路线图

- [x] 基于 azure API proxy
- [x] SSE
- [x] 用户管理
  - [x] 初始化用户管理
- [x] 基于 token 计数
- [x] Docker
- [ ] 完善 README

## 入门指南

要开始使用此项目，请按照以下步骤操作：

1. 安装 docker
2. 新建文件夹，创建配置文件
3. 选择合适方式启动

## Docker

### 安装 Docker

以下是安装 Docker 的命令行：

```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
```

### 腾讯云轻量云主机 OpenCloudOS (未安装 Docker) 的安装方法

1. 添加 Docker CE 软件源

```bash
dnf config-manager --add-repo=http://mirrors.tencent.com/docker-ce/linux/centos/docker-ce.repo
```

2. 列出可用的 Docker CE 版本

```bash
dnf list docker-ce
```

3. 安装最适合的 Docker CE 版本

```bash
dnf install -y docker-ce --nobest
```

4. 启动 Docker 服务

```bash
systemctl start docker
```

执行以上步骤后，您已经安装并启动了 Docker。现在，您可以通过以下方式启动服务。

### 构建

```shell
docker build -t hermanz/chat_azure .
```

```shell
docker buildx build --platform 'linux/arm64' hermanz/chat_azure:arm64 .
docker buildx build --platform 'linux/amd64' hermanz/chat_azure:amd64 .
```

### 部署-数据库持久存储

```shell
docker volume create chat_azure.db
```

### 部署（使用 env）

```shell
docker run -d --name chat_azure \
    -v  chat_azure.db:/data:rw \
    -e "RESOURCENAME=[AZURE's resource name]" \
    -e "APIKEY=[AZURE's api key]" \
    -e "MAPPER_GPT35TUBER=[AZURE's OpenAI deploy]" \
    -p 8080:8080 \
    hermanz/chat_azure:latest
```

### 部署（使用配置）

```shell
docker run -d --name chat_azure \
    -v  chat_azure.db:/data:rw \
    -v config.yaml:/opt/config.yaml:ro \
    -p 8080:8080 \
    hermanz/chat_azure:latest
```

## 管理 API

此项目还提供两个管理 API。

1. 添加新用户：

   当程序第一次执行时，它将创建“users”表并生成“root”用户以及其“token”（作为“admin_token”），它将打印在系统输出中，可以通过 docker logs -f chat_azure 来查看 root key。

   请求：

   ```
   curl -d '{"admin_token":"093E5AqE","username":"pig","token":"90092700"}' http://127.0.0.1:8080/v1/adduser
   ```

   响应：

   ```
       {"status":"ok"}
   ```

2. 查询用户信息：

   请求：

   ```
   curl -d '{"token":"093E5AqE"}' http://127.0.0.1:8080/v1/queryuser
   ```

   响应：

   ```
   {"count":239,"status":2,"username":"root"}
   ```

   其中，“count”表示在此程序中使用的消耗的令牌数量，用于计算使用情况。

## 许可证

该项目基于 [MIT 许可证](https://opensource.org/licenses/MIT) 授权。
