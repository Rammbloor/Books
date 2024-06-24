package methods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Err string

func (e Err) Error() string {
	return fmt.Sprintf("Ошибка чтения файла - %v ", e)
}

func newBook() string {
	fmt.Println("Введите название книги")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := "|Book title| - " + sc.Text()
	fmt.Println("Введите Автора книги")
	scAuthor := bufio.NewScanner(os.Stdin)
	scAuthor.Scan()
	author := " |Author| - " + scAuthor.Text()
	fmt.Println("Введите количество страниц в книге")
	scNumber := bufio.NewScanner(os.Stdin)
	scNumber.Scan()
	sumOfPages := " |Number of pages| - " + scNumber.Text()

	newBook := name + author + sumOfPages
	return newBook
}

func UpdateBook() {
	bookBase, err := os.ReadFile("Book-base.txt")
	if err != nil { //
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	// вывод всех книг + скан номера нужной книги
	book := strings.Split(string(bookBase), "\n")

	fmt.Println("Выберите номер книги для перезаписи данных и введите в консоль!")
	for i, v := range book {
		fmt.Printf("%d: %s\n", i, v)
	}
	// Считываем с консоли номер строчки книги
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	bookNumber, _ := strconv.Atoi(scan.Text())

	if bookNumber < 0 || bookNumber > len(book)-1 {
		fmt.Println("Введите корректный номер")

	} else {
		// вводим новые данные
		fmt.Println("Введите данные новой книги в следующем формате: ")

		// затычка временным файлом
		tmpFile, err := os.Create("temp_file.txt")
		if err != nil {
			fmt.Println("Ошибка создания временного файла:", err)
			return
		}

		//  переписываем старые данные вов временный файл
		for i, line := range book {
			if i == bookNumber {

				line = strconv.Itoa(i) + newBook()
			}
			_, err := tmpFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Ошибка записи во временный файл:", err)
				return
			}
		}

		tmpFile.Close()

		// Удаляем оригинальный файл
		err = os.Remove("Book-base.txt")
		if err != nil {
			fmt.Println("Ошибка удаления оригинального файла:", err)
			return
		}

		// Переименовываем временный файл в оригинальный
		err = os.Rename("temp_file.txt", "Book-base.txt")
		if err != nil {
			fmt.Println("Ошибка переименования временного файла:", err)
			return
		}
	}
}
