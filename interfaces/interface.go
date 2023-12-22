package interfaces

import (
	"book-management/models"
)

type DataSaver interface {
	SaveDataToCSV(fileName string, books []models.Book) error
}

type DataLoader interface {
	LoadDataFromCSV(fileName string) error
}
