# chat_azure --- Aure Open AI API Transparent Proxy

有用户管理功能的 Azure Open AI API 透明代理后台

This project is a transparent proxy for the Azure Open AI API, with user management functionality. It is based on the following repositories:

- [libli/chat](https://github.com/libli/chat)
- [haibbo/cf-openai-azure-proxy](https://github.com/haibbo/cf-openai-azure-proxy)

## Features

- User authentication and management
- Transparent proxy for the Azure Open AI API

## Roudmap

- [x] 基于 azure API proxy
- [x] SSE
- [x] 用户管理
- [x] 基于 token 计数
- [x] Docker
- [ ] 完善 README

## Getting Started

To get started with this project, follow these steps:

1. Clone this repository
2. Install the required dependencies
3. Configure the application
4. Start the server

## Docker

### Build

```shell
docker build -t chat_azure .

# optional: multi-arch
docker buildx build --platform 'linux/amd64' -t chat_azure:amd64 .
docker buildx build --platform 'linux/arm64' -t chat_azure:arm64 .
```

### Deploy (with env)

```shell
docker run -d --name chat_azure \
    -e "RESOURCENAME=[AZURE's resource name]" \
    -e "APIKEY=[AZURE's api key]" \
    -e "MAPPER_GPT35TUBER=[AZURE's OpenAI deploy]" \
    -p 8080:8080 \
    chat_azure
```

### Deploy (with config)

```shell
docker run -d --name chat_azure \
    -v config.yaml:/opt/config.yaml:ro \
    -p 8080:8080 \
    chat_azure
```

## Management API

This project also provides two management APIs.

1. Adding a new user:

   When the program is executed for the first time, it will create the `users` table and generate the `root` user along with their `token` (as `admin_token`), It will be printed in the system output.

   Request:

   ```
   curl -d '{"admin_token":"093E5AqE","username":"pig","token":"90092700"}' http://127.0.0.1:8080/v1/adduser
   ```

   Response:

   ```
       {"status":"ok"}
   ```

2. Querying user information:

   Request:

   ```
   curl -d '{"token":"093E5AqE"}' http://127.0.0.1:8080/v1/queryuser
   ```

   Response:

   ```
   {"count":239,"status":2,"username":"root"}
   ```

   Among them, "count" represents the number of tokens consumed in this program, which is used to calculate the usage.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
