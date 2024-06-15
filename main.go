package main

import (
	"bufio"
	"data-project/functions"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Cписок функций:\n" +
		"ShowAll - Показывает все книги в хранилище\n" +
		"AddBook - Внести данные про новую книгу\n" +
		"UpdateBook - Обновить по названию книги\n" +
		"DeleteBook - Удалить книгу по названию\n" +
		"Введите исполняемую функцию!")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	function := sc.Text()

	functions.FunctionSelection(function)

}
