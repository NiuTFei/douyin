# 项目介绍

该项目为字节跳动第五届青训营后端场项目作业，主要完成一个极简版抖音的后台开发。

项目API文档：https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523

# 技术架构

- 开发环境：`Golang 1.19.2`
- 数据库：`MySQL 8.0.32`、`Redis 7.0.5`
- HTTP框架：[Gin](https://github.com/gin-gonic/gin)
- ORM框架：[GORM](https://github.com/go-gorm/gorm)
- Redis客户端：[go-redis](https://github.com/redis/go-redis)

# 运行

```shell
go build main.go router.go & ./main
```

