package main

/*Requests*/

import (
	"encoding/json"
	"net/http"
	"path"
)

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		dataJson, _ := json.Marshal(books)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	dataJson, err := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)
	books = append(books, book)
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	} else {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		reqBook := Book{}
		json.Unmarshal(body, &reqBook)
		if reqBook.Id != "" { books[i].Author = reqBook.Id }
		if reqBook.Edition != "" { books[i].Author = reqBook.Edition }
		if reqBook.Author    != "" { books[i].Author = reqBook.Author }
		if reqBook.Copyright != "" { books[i].Author = reqBook.Copyright }
		if reqBook.Pages != "" { books[i].Author = reqBook.Pages }
		if reqBook.Publisher != "" { books[i].Author = reqBook.Publisher }
		if reqBook.Title != "" { books[i].Author = reqBook.Title }
		if reqBook.Language != "" { books[i].Author = reqBook.Language }
		updateBook(id, reqBook)
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	} else {
		deleteBook(id)
	}
	w.WriteHeader(200)
	return
}