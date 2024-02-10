package main

import (
	"context"
	"testing"
	"time"

	pb "blog/proto/blog/proto" 
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func TestBlogService(t *testing.T) {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewBlogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test CreatePost
	createdPost, err := client.CreatePost(ctx, &pb.Post{
		Title:           "Sample Post",
		Content:         "This is a sample post content",
		Author:          "John Doe",
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            []string{"golang", "grpc"},
	})
	if err != nil {
		t.Fatalf("could not create post: %v", err)
	}

	// Test ReadPost
	readPost, err := client.ReadPost(ctx, &pb.PostIdRequest{PostId: createdPost.PostId})
	if err != nil {
		t.Fatalf("could not read post: %v", err)
	}
	if readPost.PostId != createdPost.PostId {
		t.Fatalf("read post: expected %s, got %s", createdPost.PostId, readPost.PostId)
	}

	// Test UpdatePost
	updatedPost, err := client.UpdatePost(ctx, &pb.Post{
		PostId:          createdPost.PostId,
		Title:           "Updated Sample Post",
		Content:         "This is an updated sample post content",
		Author:          "John Doe",
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            []string{"golang", "grpc", "update"},
	})
	if err != nil {
		t.Fatalf("could not update post: %v", err)
	}
	if updatedPost.Title != "Updated Sample Post" {
		t.Fatalf("update post: expected 'Updated Sample Post', got %s", updatedPost.Title)
	}

	// Test DeletePost
	_, err = client.DeletePost(ctx, &pb.PostIdRequest{PostId: createdPost.PostId})
	if err != nil {
		t.Fatalf("could not delete post: %v", err)
	}
}
