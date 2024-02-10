# Blog

Blog is a hypothetical gRPC-based blogging platform in Go

## Project structure

```
.
├── README.md
├── client
│   └── client.go
├── go.mod
├── go.sum
├── proto
│   ├── blog
│   │   └── proto
│   │       ├── blog.pb.go
│   │       └── blog_grpc.pb.go
│   └── blog.proto
└── server
    └── server.go
```


**client/**: This directory contains all files related to the gRPC client.
**client.go**: This is the main file for the gRPC client.

**server/**: This directory contains all files related to the gRPC server.
**server.go**: This is the main file for the gRPC server.

**proto/**: This directory contains the protobuf definition file.
**blog.pb.go**: This file contains the generated Go code from the protobuf definition file blog.proto.
**blog.proto**: This is the protobuf definition file for your API.
go.mod: This file is used for Go module management. It tracks the dependencies of your project.

## How to run?
Compile the protobuf file (blog.proto) to generate the Go code (blog.pb.go) 

```
protoc --go_out=. --go-grpc_out=. blog.proto
```

Then run server.go and client.go and different terminals

## How to test?

Ensure that the gRPC server is running before executing the test cases. You can run the test cases using go test. This will execute the test cases and report any failures or errors encountered during the test execution.