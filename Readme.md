# Example typescript grpc client

Showcases how to create a grpc typescript client.

## Getting started

You'll need docker and docker-compose installed. 

Run `docker compose up`

### What happened

A go grpc server is started, after that a node js client makes a request to the grpc server. 

You should see the following output in your terminal

```shell
example-proto-typescript-dummy_backend-1  | 2023/04/13 12:37:30 I got a message request: request
example-proto-typescript-client-1         | GetExampleMessageResponse {
example-proto-typescript-client-1         |   exampleMessage: ExampleMessage { number: 42, concatString: '42_example' }
example-proto-typescript-client-1         | }
```

## Keep in mind

`client/client.ts` is configured for running in a docker container, for that reason the base url is set to 
`"http://dummy_backend:8080"`, if you want to test this script locally change that to `http://localhost:38080"`

You'll also need to generate the typescript stubs, from the root of this directory run:
```shell
npm install @bufbuild/buf @bufbuild/protoc-gen-es @bufbuild/protobuf @bufbuild/protoc-gen-connect-es @bufbuild/connect
```
to install the required dependencies and then run 
```shell
make generate-ts
```
to generate the stubs


To interact with the grpc server use [grpcurl](https://github.com/fullstorydev/grpcurl)

Some example commands:
```shell
grpcurl -plaintext localhost:38080 describe
grpcurl -plaintext -d '{"ran_number":2}' localhost:38080  example.v1.ExampleService.GetExampleMessage
```