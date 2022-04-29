## 项目结构

```
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── api
│   ├── merchant
│   │   ├── merchant.pb.go
│   │   ├── merchant.proto
│   │   ├── merchant_grpc.pb.go
│   │   └── merchant_http.pb.go
│   └── user
│       ├── user.pb.go
│       ├── user.proto
│       ├── user_grpc.pb.go
│       └── user_http.pb.go
├── cmd
│   └── admin
│       └── main.go
├── configs
│   └── config.yaml
├── go.mod
├── go.sum
├── internal
│   ├── conf
│   │   ├── conf.pb.go
│   │   └── conf.proto
│   ├── merchant
│   │   ├── biz
│   │   │   ├── businessLine.go
│   │   │   └── merchant.go
│   │   ├── data
│   │   │   ├── businessLine.go
│   │   │   └── merchant.go
│   │   └── service
│   │       ├── businessLine.go
│   │       ├── merchant.go
│   │       └── service.go
│   ├── pkg
│   │   └── data.go
│   └── user
│       ├── biz
│       │   └── user.go
│       ├── data
│       │   └── user.go
│       └── service
│           ├── service.go
│           └── user.go
├── openapi.yaml
└── third_party
    ├── README.md
    ├── errors
    │   └── errors.proto
    ├── google
    │   ├── api
    │   │   ├── annotations.proto
    │   │   ├── client.proto
    │   │   ├── field_behavior.proto
    │   │   ├── http.proto
    │   │   └── httpbody.proto
    │   └── protobuf
    │       └── descriptor.proto
    └── validate
        ├── README.md
        └── validate.proto
```



| 文件夹                    | 说明                                                     |      |
| ------------------------- | -------------------------------------------------------- | ---- |
| api                       | 存放proto文件                                            |      |
| --merchant                | merchant模块                                             |      |
| cmd                       | 存放项目启动脚本，用于项目初始化等操作                   |      |
| --admin                   | admin项目                                                |      |
| configs                   | 存放配置文件                                             |      |
| internal                  | 存放项目私有文件，不可被引用                             |      |
| --internal/user           | 用户模块（服务）                                         |      |
| ----internal/user/service | 类比ddd的application层，协同各类biz的交互                |      |
| ----internal/user/biz     | 类比ddd的domain层，处理主要的业务逻辑                    |      |
| ----internal/user/data    | 类比ddd的repo层，最终由该层来请求数据                    |      |
| --internal/merchant       | 商户模块（服务）                                         |      |
| --internal/pkg            | 存放多模块（服务）共享数据，比如数据库，消息队列的连接等 |      |
| --internal/conf           | 存放多模块（服务）共享配置结构                           |      |
| third_party               | 第三方包                                                 |      |

