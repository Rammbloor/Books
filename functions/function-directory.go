package functions

import (
	"data-project/methods"
	"fmt"
)

func FunctionSelection(s string) {

	switch s {
	case "1":
		methods.ShowAll()
	case "2":
		methods.ShowBookToNumber()
	case "3":
		var NewBook methods.Books
		NewBook.AddBook()
	case "4":
		methods.UpdateBook()
	case "5":
		methods.DeleteBook()

	default:
		fmt.Println("Введите корректный номер!")

	}

}
