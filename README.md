## 设计思想和原则

框架中用到的设计思想和原则，尽量满足 "高内聚、低耦合"，主要遵从下面几个原则
- 1. 单一职责原则
- 2. 基于接口而非实现编程
- 3. 依赖注入
- 4. 多用组合
- 5. 迪米特法则
> 迪米特法则: 不该有直接依赖关系的类之间，不要有依赖；有依赖关系的类之间，尽量只依赖必要的接口


## ✨ 技术栈

- 框架路由使用 [Gin](https://github.com/gin-gonic/gin) 路由
- 中间件使用 [Gin](https://github.com/gin-gonic/gin) 框架的中间件
- 数据库组件 [GORM](https://github.com/jinzhu/gorm)
- 文档使用 [Swagger](https://swagger.io/) 生成
- 配置文件解析库 [Viper](https://github.com/spf13/viper)
- 使用 [JWT](https://jwt.io/) 进行身份鉴权认证
- 校验器使用 [validator](https://github.com/go-playground/validator)  也是 Gin 框架默认的校验器
- 任务调度 [cron](https://github.com/robfig/cron)
- 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- 测试框架 [GoConvey](http://goconvey.co/)
- CI/CD [GitHub Actions](https://github.com/actions)
- 使用 [GolangCI-lint](https://golangci.com/) 进行代码检测
- 使用 make 来管理 Go 工程
- 使用 shell(admin.sh) 脚本来管理进程
- 使用 YAML 文件进行多环境配置

## 📗 目录结构

```shell
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
目录的定义与规范：
internal 目录下的包，不允许被其他项目中进行导入；项目内部进行共享的包，而不期望外部共享，可以放到/internal/pkg
pkg 放置可以被外部程序安全导入的包
api grpc客户端和Swagger文档，以及他们生成的文件

参考资料：项目目录结构 https://lailin.xyz/post/go-training-week4-project-layout.html

## 💻 常用命令

- make help 查看帮助
- make dep 下载 Go 依赖包
- make build 编译项目
- make gen-docs 生成接口文档
- make test-coverage 生成测试覆盖
- make lint 检查代码规范


## 🚀 部署

### 单独部署

上传到服务器后，直接运行命令即可

```bash
./scripts/admin.sh start
```