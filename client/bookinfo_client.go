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
	localDefaultAddr := "localhost:8080"
	address := os.Getenv("ADDRESS")
	if address == ""{
		address = localDefaultAddr
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddBook(ctx, &pb.Book{
		Id:        "100",
		Title:     "Sample Book",
		Edition:   "1st",
		Copyright: "2020",
		Language:  "ENGLISH",
		Pages:     "420",
		Author:    "Gustavo Adolfo Marquez",
		Publisher: "Downtown Publisher"})
	if err != nil {
		log.Fatalf("Could not add book: %v", err)
	}

	log.Printf("Book ID: %s added successfully", r.Value)

	book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get book: %v", err)
	}
	log.Printf("Book: ", book.String())

	//Update
	u, err2 := c.UpdateBook(ctx, book)
	if err2 != nil {
		log.Fatalf("Could not update book: %v", err2)
	}
	log.Printf("Book ID: %s updated successfully", u.Value)

	//delete
	book1, err2 := c.DeleteBook(ctx, &pb.BookID{Value: r.Value})
	if err2 != nil {
		log.Fatalf("Could not get book: %v", err2)
	}
	log.Printf("Deleted Book: ", book1.String())
	
	readData("books.csv")

	for i, _ := range books {
		index:=i
		if index == 0 {
			index = 1
		}
		indexBook := books[index]
		r, err := c.AddBook(ctx, &pb.Book{
			Id:        indexBook.Id,
			Title:     indexBook.Title,
			Edition:   indexBook.Edition,
			Copyright: indexBook.Copyright,
			Language:  indexBook.Language,
			Pages:     indexBook.Pages,
			Author:    indexBook.Author,
			Publisher: indexBook.Publisher})
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

	// added csv books and showed them
}
