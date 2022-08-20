<img src="./go.png" width="80" height="80" alt="logo">

# gRPC Communication with go

### About

> My study of the full cycle course module of grpc with go lang

### Concepts

- gRPC
- Protobuf
- Streaming

### Commands

install
```bash
sudo apt install protobuf-compiler
```
```bash
brew install protobuf
```
```bash
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

```
Generate stubs
```bash
protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go_out=.
protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=.
```
Run server
```bash
go run cmd/server/server.go
 
go run cmd/client/client.go 
```
