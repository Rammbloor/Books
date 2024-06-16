package main

import (
	"bufio"
	"data-project/functions"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Cписок функций:\n" +
		"1-ShowAll - Показывает все книги в хранилище\n" +
		"2-ShowBookToNumber - Показывает книгу по введеному числу\n" +
		"3-AddBook - Внести данные про новую книгу\n" +
		"4-UpdateBook - Обновить книгу\n" +
		"5-DeleteBook - Удалить книгу\n" +
		"Введите исполняемую функцию!")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	function := sc.Text()

	functions.FunctionSelection(function)

}
