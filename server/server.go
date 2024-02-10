package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	pb "blog/proto/blog/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// server struct holds the data and implements the gRPC server interface
type server struct {
	posts map[string]*pb.Post // Map to store posts (simulating a simple in-memory database)
	mu    sync.RWMutex        // Mutex for thread safety
	pb.UnimplementedBlogServiceServer 
}

// CreatePost creates a new blog post
func (s *server) CreatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate a unique PostID (You might want to use a better approach in production)
	req.PostId = generatePostID()

	// Add the post to the map
	s.posts[req.PostId] = req

	log.Printf("Post created: %+v", req)
	return req, nil
}

// ReadPost retrieves a blog post by PostID
func (s *server) ReadPost(ctx context.Context, req *pb.PostIdRequest) (*pb.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	post, ok := s.posts[req.PostId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Post with ID %s not found", req.PostId)
	}

	log.Printf("Post retrieved: %+v", post)
	return post, nil
}

// UpdatePost updates an existing blog post
func (s *server) UpdatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the post exists
	_, ok := s.posts[req.PostId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Post with ID %s not found", req.PostId)
	}

	// Update the post in the map
	s.posts[req.PostId] = req

	log.Printf("Post updated: %+v", req)
	return req, nil
}

// DeletePost deletes a blog post by PostID
func (s *server) DeletePost(ctx context.Context, req *pb.PostIdRequest) (*pb.DeleteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the post exists
	_, ok := s.posts[req.PostId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Post with ID %s not found", req.PostId)
	}

	// Delete the post from the map
	delete(s.posts, req.PostId)

	log.Printf("Post deleted: %s", req.PostId)
	return &pb.DeleteResponse{Success: true}, nil
}

// generatePostID generates a unique PostID (just a simple example)
func generatePostID() string {
	// You might want to use a more robust method to generate PostIDs in production
	return "post" + time.Now().Format("20060102150405")
}

func main() {
	// Initialize an empty map to store posts
	posts := make(map[string]*pb.Post)

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the server implementation to handle gRPC requests
	pb.RegisterBlogServiceServer(s, &server{posts: posts})

	// Listen for incoming connections on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the gRPC server
	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
