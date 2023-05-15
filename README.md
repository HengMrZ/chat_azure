# chat_azure --- Aure Open AI API Transparent Proxy

有用户管理功能的Azure Open AI  API透明代理后台

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
- [ ] Docker
- [ ] 完善 README

## Getting Started

To get started with this project, follow these steps:

1. Clone this repository
2. Install the required dependencies
3. Configure the application
4. Start the server

## Management API

This project also provides two management APIs.

1. Adding a new user:

    Request:
    ```
    curl -d '{"admin_token":"093E5AqE","username":"pig","token":"90092700"}' http://127.0.0.1:3389/v1/adduser
    ```

    Response:
    ```
        {"status":"ok"}
    ```
2. Querying user information:

    Request:
    ```
    curl -d '{"token":"093E5AqE"}' http://127.0.0.1:3389/v1/queryuser
    ```
    Response:
    ```
    {"count":239,"status":2,"username":"root"}
    ```

    Among them, "count" represents the number of tokens consumed in this program, which is used to calculate the usage.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

