**# tiktok**

大家好，这里是艺画开天队的Tiktok大项目

这里是项目的总目录（根目录）



**# 第一次Clone tiktok项目时，您需要做什么？**

构建镜像并启动服务

```shell
docker-compose up -d 
```

启动 `user` 服务

```shell
# 进入 user 目录
cd user
go run .
```

启动 `video`服务

```shell
# 进入 video 目录
cd video
go run .
```

启动 `composite` 服务

```shell
# 进入 composite 目录
cd composite
go run .
```

启动 `gateway`

```shell
# 进入 gateway 目录
cd gateway
go run .
```

