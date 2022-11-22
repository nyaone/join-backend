# 喵窝 邀请管理系统 - 后端

## 部署

### 准备

1. 您需要拥有至少以下两个账号：
   1. 具有实例 `管理员` 身份的账号，用于实际操作新用户的注册。注意不能仅有 `监察员` 身份，需要是实例管理员级别。
   2. 一个一般账号，用于发送登录使用的验证消息和迎新时候的消息。出于安全因素考虑，最好仅是普通用户，没有任何管理权限。
2. 生成 API Token （设置 - 其他设置 - API - 生成访问令牌）
   1. 对管理员身份的访问令牌，不需要开启任何权限选项，直接生成就可以。
   2. 对一般账号的访问令牌，要求开启 `撰写或删除消息` 和 `撰写或删除帖子` 两项权限。

### 使用 docker 与 docker-compose 部署

1. 在您的服务器上新建目录，复制本项目仓库里的 `docker-compose.yml` 和 `config.yml.example` 。
2. 将 `config.yml.example`  重命名为 `config.yml` ，并编辑其中的各项参数，参数对应的解释均已注释在配置文件中。
3. 使用 `docker-compose pull` 拉取预构建的镜像。
4. 使用 `docker-compose up -d` 启动程式组合。

### 使用二进制方式部署

请自行构建相关文件，本程式使用 Go 语言开发，您可以参照 Dockerfile 文件中指定的构建指令进行构建。
