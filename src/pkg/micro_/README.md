
安装protobuf库文件

go get github.com/golang/protobuf/proto

安装插件

go get github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/protoc-gen-micro/v2

protoc --proto_path=src/pkg/micro_/protos --micro_out=src/pkg/micro_/protos  --go_out=src/pkg/micro_/protos Users.proto

protoc --proto_path=protos --micro_out=protos  --go_out=protos Hello.proto

micro new greeter



### 注意

github.com/micro/go-micro/v2

v2 版本只能在1.15下使用


### 巨坑的玩意

MICRO_REGISTRY_ADDRESS 不能为a.b.c 这种，不知道为什么，换成下面就好了


https://github.com/microhq/examples/tree/master/api/api
```
Run the micro API with custom namespace

micro api --handler=api --namespace=com.foobar.api
or

MICRO_API_NAMESPACE=com.foobar.api micro api --handler=api
Set service name with the namespace

service := micro.NewService(
        micro.Name("com.foobar.api.example"),
)

```
