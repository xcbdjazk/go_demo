
安装protobuf库文件

go get github.com/golang/protobuf/proto

安装插件

go get github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/protoc-gen-micro/v2

protoc --proto_path=src/pkg/micro_/protos  --go_out=src/pkg/micro_/Users Users.proto



### 注意

github.com/micro/go-micro/v2

v2 版本只能在1.15下使用