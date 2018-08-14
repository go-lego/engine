# engine


### Generate Protobuf
```
protoc --proto_path=$GOPATH/src:eds/proto --go_out=eds/proto --micro_out=eds/proto eds/proto/*.proto
protoc --proto_path=$GOPATH/src:proto --go_out=proto --micro_out=proto proto/*.proto
```