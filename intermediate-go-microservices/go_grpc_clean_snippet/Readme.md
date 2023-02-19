> [Short video explaination of RPC/GRPC Orientations](https://murf.ai/share/le2t77ym%20)

## Generating pb.go & grpc.pb.go files via .proto file

```bash

# install protoc
choco install protoc
protoc --version
OR
# reff: https://grpc.io/docs/languages/go/quickstart/ & install and setup protobuf compiler https://protobuf.dev/downloads/

# generate pb.go & grpc.pb.go files
# cd to go_grpc_clean_snippet
protoc --go_out=. --go-grpc_out=. proto/greet.proto

# manage depend
go mod tidy
```
