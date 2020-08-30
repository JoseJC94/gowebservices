package main

import (
	"context"
	pb "github.com/JoseJC94/gowebservices/booksapp"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	address := os.Getenv("ADDRESS")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	readData("books.csv")

	for _, currentBook := range books {
		r, err := c.AddBook(ctx, &pb.Book{
			Id:        currentBook.Id,
			Title:     currentBook.Title,
			Edition:   currentBook.Edition,
			Copyright: currentBook.Copyright,
			Language:  currentBook.Language,
			Pages:     currentBook.Pages,
			Author:    currentBook.Author,
			Publisher: currentBook.Publisher})
		if err != nil {
			log.Fatalf("Could not add book: %v", err)
		}

		log.Printf("Book ID: %s added successfully", r.Value)
		book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
		if err != nil {
			log.Fatalf("Could not get book: %v", err)
		}
		log.Printf("Book: ", book.String())
	}

}
