## rpc

RPC client 接口定义层，基于 protobuf 严格定义 RPC 接口路由、参数和文档。

## 准备

### 安装 protoc 编译器

查看版本

```bash
protoc --version
libprotoc 3.12.1
```

### 安装 protoc-gen-go 插件

运行：

```shell script
go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

### 安装 grpc-go

grpc-go包含了Go的grpc库

```shell script
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### 安装 twirp
它的设置非常简单，这对我来说最重要；
同时支持 http 1.1 和 http 2.0；
同时支持 Protobuf 和 JSON；
易于调试

```shell script
go get github.com/twitchtv/twirp/protoc-gen-twirp
go install github.com/twitchtv/twirp/protoc-gen-twirp
```

### 编辑 proto 文件

### 生成代码

//生成grpc文件
protoc --go-grpc_out=.  user.proto

//生成一般文件
protoc --go_out=.  user.proto

//生成twirp文件
protoc --twirp_out=.  user.proto

-I：源文件的目录（可省略）
--go_out: 设置所生成的Go代码输出目录

## 目录结构

通常一个服务一个文件夹。服务下有版本，一个版本一个文件夹。内部服务一般使用 v0 作为版本。

一个版本可以定义多个 service，每个 service 一个 proto 文件。

典型的目录结构如下：

```bash
rpc/user # 业务服务
└── v0   # 服务版本
    ├── user.pb.go     # protobuf message 定义代码[自动生成]
    └── user.proto     # protobuf 描述文件[业务方定义]
```
1、
## 定义接口

服务接口使用 protobuf 描述。

```proto
syntax = "proto3";

package user.v0; // 包名，与目录保持一致

// 服务名，只要能定义一个 service
service Echo {
  // 服务方法，按需定义
  rpc Hello(HelloRequest) returns (HelloResponse);
}

// 入参定义
message HelloRequest {
  // 字段定义，如果使用 form 表单传输，则只支持
  // int32, int64, uint32, unint64, double, float, bool, string,
  // 以及对应的 repeated 类型，不支持 map 和 message 类型！
  // 框架会自动解析并转换参数类型
  // 如果用 json 或 protobuf 传输则没有限制
  string message = 1; // 这是行尾注释，业务方一般不要使用
  int32 age = 2;
  // form 表单格式只能部分支持 repeated 语义
  // 但客户端需要发送英文逗号分割的字符串
  // 如 ids=1,2,3 将会解析为 []int32{1,2,3}
  repeated int32 ids = 3;
}

message HelloMessage {
  string message = 1;
}

// 出参定义,
// 理论上可以输出任意消息
// 但我们的业务要求只能包含 code, msg, data 三个字段，
// 其中 data 需要定义成 message
// 开源版本可以怱略这一约定
message HelloResponse {
  // 业务错误码[机读]，必须大于零
  // 小于零的主站框架在用，注意避让。
  int32 code = 1;
  // 业务错误信息[人读]
  string msg = 2;
  // 业务数据对象
  HelloMessage data = 3;
}
```

2、
```
# 针对指定服务
protoc --go_out=. --twirp_out=. echo.proto

# 针对所有服务
find rpc -name '*.proto' -exec protoc --twirp_out=. --go_out=. {} \;
```

生成的文件中 *.pb.go 是由 protobuf 消息的定义代码，同时支持 protobuf 和 json。*.twirp.go 则是 rpc 路由相关代码。

## 实现接口

请参考 [server/README.md](https://github.com/go-eagle/eagle/tree/master/internal/server/README.md)

## 自动注册

eagle 提供的脚手架可以自动生成 proto 模版、server 模版，并注册路由。 运行以下命令：

```bash
go run cmd/eagle/main.go rpc --server=foo --service=echo
```

3、
会自动生成

```bash
rpc
└── foo
    └── v1
        ├── echo.pb.go
        ├── echo.proto
        └── echo.twirp.go
server
└── fooserver1
    └── echo.go
```

## Reference

- [Proto 代码到底放哪里？](https://eddycjy.com/posts/where-is-proto/)
- [Go 每日一库之 twirp](https://zhuanlan.zhihu.com/p/147391981)