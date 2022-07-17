# 项目说明
lebron为工程名，lebron下面有apps和pkg两个目录，其中apps存放的是我们所有的微服务，比如order为订单相关的微服务，pkg目录为所有服务共同依赖的包的存放路径，比如所有的服务都需要依赖鉴权就可以放到pkg目录下。

- app - BFF服务
- cart - 购物车服务
- order - 订单服务
- pay - 支付服务
- product - 商品服务
- recommend - 推荐服务
- reply - 评论服务
- user - 账号服务

在每个服务目录下我们又会分为多个服务，主要会有如下几类服务：

- api - 对外的BFF服务，接受来自客户端的请求，暴露HTTP接口
- rpc - 对内的微服务，仅接受来自内部其他微服务或者BFF的请求，暴露gRPC接口
- rmq - 负责进行流式任务处理，上游一般依赖消息队列，比如kafka等
- admin - 也是对内的服务，区别于rpc，更多的是面向运营侧的且数据权限较高，通过隔离可带来更好的代码级别的安全，直接提供HTTP接口

# 使用说明
## rpc代码生成
执行如下命令即可初始化order rpc代码
```sh
cd lebron/apps/order
goctl rpc new rpc
```
## api代码生成
执行如下命令即可初始化order admin代码，注意order admin为`api服务`，直接对前端提供HTTP接口
```sh
goctl api new admin
```

# RPC定义

因为BFF只负责数据的组装工作，数据真正的来源是各个微服务通过RPC接口提供，接下来我们来定义各个微服务的proto
order.proto
```protobuf
syntax = "proto3";

package order;
option go_package="./order";


service Order {
  rpc Orders(OrdersRequest) returns(OrdersResponse);
}

message OrdersRequest {
  int64 user_id = 1;
  int32 status = 2;
  int64 cursor = 3;
  int32 ps = 4;
}

message OrdersResponse {
  repeated OrderItem orders = 1;
  bool is_end = 2;
  string create_time = 3;
}

message OrderItem {
  string order_id = 1;
  int64 quantity = 2;
  float payment = 3;
  int64 product_id = 4;
  int64 user_id = 5;
  int64 create_time = 6;
}
```
使用如下命令重新生成代码，注意这里需要依赖protoc-gen-go和protoc-gen-go-grpc两个插件，木有安装的话执行下面命令会报错
```shell
goctl rpc protoc order.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```