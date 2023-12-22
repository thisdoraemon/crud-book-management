package helper

import (
	"book-management/models"
	"book-management/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CSVDataSaver struct{}

func (saver *CSVDataSaver) SaveDataToCSV(fileName string, books []models.Book) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	for _, book := range books {
		row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author +
			"," + book.ReleaseYear + "," + strconv.Itoa(book.Pages) + "\n"

		file.WriteString(row)
	}
	return nil
}

type CSVDataLoader struct{}

// LoadDataFromCSV loads data from a CSV file
func (loader *CSVDataLoader) LoadDataFromCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	booksService := &service.BookService{}

	for scanner.Scan() {
		record := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(record[0])
		pages, _ := strconv.Atoi(record[4])

		book := models.Book{
			Id:          id,
			Title:       record[1],
			Author:      record[2],
			ReleaseYear: record[3],
			Pages:       pages,
		}
		booksService.Books = append(booksService.Books, book)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	return nil
}

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("File", fileName, "berhasil dibuat")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)
}
