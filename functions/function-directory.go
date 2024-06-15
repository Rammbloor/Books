package functions

import "data-project/methods"

func FunctionSelection(s string) {

	switch s {

	case "ShowAll":
		methods.ShowAll()
	case "AddBook":
		var NewBook methods.Books
		NewBook.AddBook()
	case "UpdateBook":
		methods.UpdateBook()
	case "DeleteBook":
		methods.DeleteBook()

	}

}
