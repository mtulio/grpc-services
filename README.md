# grpc-services

gRPC Services exposing protobuf interfaces in Golang.

That project is used to study purpose that implements some usefull services
interface using gRPC in Golang.

## crawler

Crawler is an web crawler.

## grpc-server

`grpc-server` is a server that implements packages created in `grpc-services` and
exposes it as a reflection.

There are two ways to call the functions exposed by the server:

1. implement an client.
1. using an CLI to call the functions.
 * [Evans CLI](https://github.com/ktr0731/evans)
* [grpc_cli](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md)

