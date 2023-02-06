> GRPC golang microservice snippet

**reff: https://grpc.io/**

- Services are defined via protocol buffers that serializes/deserializes binary data flows
- since the approach is oriented towards binary data grpc is highly performance oriented and outperforms restfull or graphql architecture styles.

> developing grpc

```bash

#proto-file
here services and method for services are defined with message formats

## typical service defination
service "YourServiceName" {
    // define your methods here....
    rpc MethodName()
}

## typical message formats, Input type i.e request message & Output type i.e response message
message ClientRequest{
    string fieldName = IndexPosition/FieldNumber inside the binary protocol;
}

```

- **IMP- Field numbers are an important part of Protobuf. They're used to identify fields in the binary encoded data, which means they can't change from version to version of your service. The advantage is that backward compatibility and forward compatibility are possible.**

- proto buffs have its own types ref: https://developers.google.com/protocol-buffers/docs/proto3

- proto lang is language and platform neutral & independent, a simple double type message format fieldName specified in .proto file is language independent reff: https://developers.google.com/protocol-buffers/docs/proto3#scalar

- In contrast to json where encoding is as string in protobuf it is binary reff: https://developers.google.com/protocol-buffers/docs/encoding#length-types

> Generating code files for specific language on basis of .proto file via protoc

- **install protobuff refer: https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/**

```bash

# open cmd/terminal
protoc --help

# download/install depend ref:https://grpc.io/docs/languages/go/quickstart/
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"


# navigate to grpc_golang_snippet
# generate code file, mention the input and the out code go file
protoc -I protos/ protos/fiat.proto --go-grpc_out=protos

```

- grpc package https://pkg.go.dev/google.golang.org/grpc maintained by google
