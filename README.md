# Blog

Blog is a hypothetical gRPC-based blogging platform in Go

## Project Description
The API supports CRUD operations for blog posts, and each post have the following attributes:

PostID (unique identifier)
Title
Content
Author
Publication Date
Tags (multiple tags per post)

The API supports the following operations:-

### Create Post
Input: Post details (Title, Content, Author, Publication Date, Tags)

Output: The Post (PostID, Title, Content, Author, Publication Date, Tags). Error message, if creation fails.

### Read Post
Input: PostID of the post to retrieve

Output: Post details (PostID, Title, Content, Author, Publication Date, Tags) or an error message if the post is not found.

### Update Post
Input: PostID of the post to update and new details (Title, Content, Author, Tags)

Output: Post details (PostID, Title, Content, Author, Publication Date, Tags) or error message if the update failed

### Delete Post
Input: PostID of the post to delete

Output: Success/Failure message

### List Posts
Input: No input

Output: List all posts or error message if listing fails

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