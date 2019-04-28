# grpc-services
gRPC Services exposing protobuf interfaces in Golang

## crawler

Crawler is an web crawler.


## grpc-server

`grpc-server` is a server that implements packages created in `grpc-services` and
exposes it as a reflection.

There are two ways to call the functions exposed by the server:
1. implement an client
1. using an CLI to call the functions. E.g.: [Evans CLI](https://github.com/ktr0731/evans)

