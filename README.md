# 目录结构

## api

存放proto文件

## cmd

存放初始化项目脚本

## configs

存放配置文件

## internal

存放项目私有文件，不可被引用

### internal/user

用户模块（服务）

#### internal/user/service

类比ddd的application层，协同各类biz的交互

#### internal/user/biz

类比ddd的domain层，处理主要的业务逻辑

在该层定义实现实体的接口，再由data层去实现对应的接口

#### internal/user/data

类比ddd的repo层，最终由该层来请求数据

### internal/merchant

商户模块（服务）

### internal/pkg

存放多模块（服务）共享数据，比如数据库，消息队列的连接等

### internal/conf

存放多模块（服务）共享配置结构

## third_party

第三方包
