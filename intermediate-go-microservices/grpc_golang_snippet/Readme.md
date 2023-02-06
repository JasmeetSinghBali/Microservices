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
