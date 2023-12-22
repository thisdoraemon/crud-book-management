package service

import (
	"book-management/interfaces"
	"book-management/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BookService struct {
	Books      []models.Book
	FileName   string
	DataLoader interfaces.DataLoader
	DataSaver  interfaces.DataSaver
}

func NewBookService(fileName string, dataLoader interfaces.DataLoader, dataSaver interfaces.DataSaver) *BookService {
	return &BookService{
		Books:      []models.Book{},
		FileName:   fileName,
		DataLoader: dataLoader,
		DataSaver:  dataSaver,
	}
}

func (s *BookService) AddNewBook() error {
	var newBook models.Book

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter book details")

	fmt.Print("Book id:")
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

	_, err := s.FindBookById(newBook.Id)
	if err != nil {
		s.Books = append(s.Books, newBook)
	} else {
		return fmt.Errorf("book with id: %d already exists", newBook.Id)
	}

	err = s.DataSaver.SaveDataToCSV(s.FileName, s.Books)
	if err != nil {
		return err
	}
	fmt.Println("Book added successfully")

	return nil
}

func (s *BookService) UpdateBook() error {
	fmt.Print("Enter Book Id to update: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	bookIndex, err := s.FindBookById(bookId)
	if err != nil {
		return err
	}

	var updatedBook models.Book
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

	s.Books[bookIndex.Id] = updatedBook
	err = s.DataSaver.SaveDataToCSV(s.FileName, s.Books)
	if err != nil {
		return err
	}
	fmt.Println("Book updated successfully")
	return nil
}

func (s *BookService) DeleteBook() error {
	fmt.Print("Enter Book Id to delete: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	bookIndex, err := s.FindBookById(bookId)
	if err != nil {
		return err
	}

	s.Books = append(s.Books[:bookIndex.Id], s.Books[bookIndex.Id+1:]...)
	err = s.DataSaver.SaveDataToCSV(s.FileName, s.Books)
	if err != nil {
		return err
	}
	fmt.Println("Book deleted successfully")
	return nil
}

func (s *BookService) ViewAllBooks() error {
	if len(s.Books) == 0 {
		return fmt.Errorf("no books available")
	}

	for i, book := range s.Books {
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

func (s *BookService) FindBookById(id int) (models.Book, error) {
	for _, book := range s.Books {
		if book.Id == id {
			return book, nil
		}
	}
	return models.Book{}, fmt.Errorf("id: %d not found", id)
}
