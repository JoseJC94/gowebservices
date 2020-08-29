package main

func find(x string) int {
	for i, book := range books {
		if x == book.Id {
			return i
		}
	}
	return -1
}

func updateBook(id string, updatingBook Book){
	deleteBook(id)
	books = append(books, updatingBook)
}

func deleteBook(id string){
	i := find(id)
	books = append(books[:i], books[i+1:]...)
}


