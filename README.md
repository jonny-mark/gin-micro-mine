## 📗 目录结构

```
├── Makefile                     # 项目管理文件
├── api                          # grpc客户端和Swagger 文档
├── common                       # 公共方法 或 常数
├── cmd                          # 脚手架目录
├── config                       # 配置文件统一存放目录
├── deploy                       # 主要存放一些部署相关的配置文件和脚本
├── docs                         # 框架相关文档
├── internal                     # 业务目录
│   ├── cache                    # 基于业务封装的cache
│   ├── handler                  # http 接口
│   ├── middleware               # 自定义中间件
│   ├── model                    # 数据库 model
│   ├── repository               # 数据访问层
│   ├── routers                  # 业务路由
│   ├── server                   # http server 和 grpc server
│   ├── service                  # 业务逻辑层
│   └── web                      # web登陆
├── log                          # 存放日志的目录
├── main.go                      # 项目入口文件
├── pkg                          # 公共的 package
├── test                         # 单元测试依赖的配置文件，主要是供docker使用的一些环境配置文件
└── scripts                      # 存放用于执行各种构建，安装，分析等操作的脚本
```