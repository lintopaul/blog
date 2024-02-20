package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "blog/proto/blog/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// // Test CreatePost
	// _, err = c.CreatePost(ctx, &pb.Post{
	// 	Title:           "Sample Post",
	// 	Content:         "This is a sample post content",
	// 	Author:          "John Doe",
	// 	PublicationDate: time.Now().Format(time.RFC3339),
	// 	Tags:            []string{"golang", "grpc"},
	// })
	// if err != nil {
	// 	log.Fatalf("could not create post: %v", err)
	// }
	// log.Println("Post created successfully")

	// Test ListPosts
	p, err := c.ListPosts(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("Failed to list posts: %v", err)
	}

	posts, err := p.Recv()
	if err != nil {
		log.Fatalf("Failed to receive posts: %v", err)
	}

	// Iterate over each post in the post list
	for _, post := range posts.Posts {
		fmt.Printf("Post ID: %s\n", post.PostId)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Content: %s\n", post.Content)
		fmt.Printf("Author: %s\n", post.Author)
		fmt.Printf("Publication Date: %s\n", post.PublicationDate)
		fmt.Printf("Tags: %v\n", post.Tags)
		fmt.Println("----------------------------------------")
	}

// 	// Test UpdatePost
// 	updatedPost, err := c.UpdatePost(ctx, &pb.Post{
// 		PostId:          "post20240210173328",
// 		Title:           "Updated Sample Post",
// 		Content:         "This is an updated sample post content",
// 		Author:          "John Doe",
// 		PublicationDate: time.Now().Format(time.RFC3339),
// 		Tags:            []string{"golang", "grpc", "update"},
// 	})
// 	if err != nil {
// 		log.Fatalf("could not update post: %v", err)
// 	}
// 	log.Printf("Updated post: %v", updatedPost)

// 	// Test DeletePost
// 	_, err = c.DeletePost(ctx, &pb.PostIdRequest{PostId: "post20240210173328"})
// 	if err != nil {
// 		log.Fatalf("could not delete post: %v", err)
// 	}
// 	log.Println("Post deleted successfully")

 }
