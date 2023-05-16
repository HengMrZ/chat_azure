# chat_azure --- Aure Open AI API Transparent Proxy

This project is a transparent proxy for the Azure Open AI API, with user management functionality. It is based on the following repositories:

- [libli/chat](https://github.com/libli/chat)
- [haibbo/cf-openai-azure-proxy](https://github.com/haibbo/cf-openai-azure-proxy)

## Features

- User authentication and management
- Transparent proxy for the Azure Open AI API

## Roudmap

- [x] Based on azure API proxy
- [x] SSE
- [x] User management
- [x] Based on token counting
- [x] Docker
- [ ] Improve README

## Getting Started

To get started with this project, follow these steps:

1. Clone this repository
2. Install the required dependencies
3. Configure the application
4. Start the server

## Docker
### Install Docker

The following is the command line to install Docker:

```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
```

### Installation of OpenCloudOS (not installed Docker) on Tencent Cloud Light Cloud Host

1. Add Docker CE software source

```bash
dnf config-manager --add-repo=http://mirrors.tencent.com/docker-ce/linux/centos/docker-ce.repo
```

2. List available Docker CE versions

```bash
dnf list docker-ce
```

3. Install the most appropriate Docker CE version

```bash
dnf install -y docker-ce --nobest
```

4. Start the Docker service

```bash
systemctl start docker
```

After performing the above steps, you have installed and started Docker. Now, you can use Git for version control to store and manage code.

### Deploy (with env)

```shell
docker run -d --name chat_azure \
    -e "RESOURCENAME=[AZURE's resource name]" \
    -e "APIKEY=[AZURE's api key]" \
    -e "MAPPER_GPT35TUBER=[AZURE's OpenAI deploy]" \
    -p 8080:8080 \
    hermanz/chat_azure:latest
```

### Deploy (with config)

```shell
docker run -d --name chat_azure \
    -v config.yaml:/opt/config.yaml:ro \
    -p 8080:8080 \
    hermanz/chat_azure:latest
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
