# grpc-demo

### 定义proto

- meta.proto
- helloworld.proto

### 生成go文件

`protoc --go_out=plugins=grpc:./go/helloworld ./helloworld.proto`
`protoc --go_out=plugins=grpc:./go/meta ./meta.proto`

### 生成PHP文件

`protoc --proto_path=./ --php_out=phpClient/meta --grpc_out=phpClient/meta --plugin=protoc-gen-grpc=./grpc_php_plugin meta.proto`


### 编写server(golang)

 `main.go`
 
 运行 `go run main.go`
### client调用

`phpClient/client.php`

执行 `php phpClient/client.php` 调用rpc服务

显示结果
`{"appId":"123","title":"my title","name":"My Name","logo":"https://xxx.com"}
 the greet: Hello world--from go server
`
