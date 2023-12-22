package main

import (
	"book-management/service"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=====================================================")
	fmt.Println("==============Welcome to Book Management=============")
	for {
		fmt.Println("Please choose an action:")
		fmt.Println("1. View All Books")
		fmt.Println("2. Add New Book")
		fmt.Println("3. Update Book")
		fmt.Println("4. Delete Book")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice (1-5): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			err := service.ViewAllBooks()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			err := service.AddNewBook()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
			err := service.UpdateBook()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 4:
			err := service.DeleteBook()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 5:
			fmt.Println("Exiting Book Management. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}
