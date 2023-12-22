package service

import (
	"book-management/helper"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear string
	Pages       int
}

var Books []Book
var FileName string = "data.csv"

func AddNewBook() error {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Book Details")

	fmt.Print("Book Id:")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Book Title : ")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author : ")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Release Year : ")
	scanner.Scan()
	newBook.ReleaseYear = scanner.Text()

	fmt.Print("Pages : ")
	scanner.Scan()
	newBook.Pages, _ = strconv.Atoi(scanner.Text())

	_, err := FindBookById(newBook.Id)
	if err != nil {
		Books = append(Books, newBook)
	} else {
		return fmt.Errorf("book with id: %d already exist", newBook.Id)
	}

	err = helper.SaveDataToCSV(FileName)
	if err != nil {
		return err
	}
	fmt.Println("Book added successfully")

	return nil
}

func UpdateBook() error {
	fmt.Print("Enter Book Id to update: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	bookIndex, err := FindBookById(bookId)
	if err != nil {
		return err
	}

	var updatedBook Book
	fmt.Println("Enter Updated Book Details:")
	fmt.Print("Book Title : ")
	scanner.Scan()
	updatedBook.Title = scanner.Text()

	fmt.Print("Book Author : ")
	scanner.Scan()
	updatedBook.Author = scanner.Text()

	fmt.Print("Release Year : ")
	scanner.Scan()
	updatedBook.ReleaseYear = scanner.Text()

	fmt.Print("Pages : ")
	scanner.Scan()
	updatedBook.Pages, _ = strconv.Atoi(scanner.Text())

	Books[bookIndex.Id] = updatedBook
	err = helper.SaveDataToCSV(FileName)
	if err != nil {
		return err
	}
	fmt.Println("Book updated successfully")
	return nil
}

func DeleteBook() error {
	fmt.Print("Enter Book Id to delete: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	bookIndex, err := FindBookById(bookId)
	if err != nil {
		return err
	}

	Books = append(Books[:bookIndex.Id], Books[bookIndex.Id+1:]...)
	err = helper.SaveDataToCSV(FileName)
	if err != nil {
		return err
	}
	fmt.Println("Book deleted successfully")
	return nil
}

func ViewAllBooks() error {
	if len(Books) == 0 {
		return fmt.Errorf("no books available")
	}

	for i, book := range Books {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("Book - ", i+1)
		fmt.Println("Book Id :", book.Id)
		fmt.Println("Book Title :", book.Title)
		fmt.Println("Book Author :", book.Author)
		fmt.Println("Release Year :", book.ReleaseYear)
		fmt.Println("Pages :", book.Pages)
		fmt.Println(strings.Repeat("=", 50))
	}
	return nil
}

func FindBookById(id int) (Book, error) {
	for _, book := range Books {
		if book.Id == id {
			return book, nil
		}
	}
	return Book{}, fmt.Errorf("id: %d not found", id)
}
