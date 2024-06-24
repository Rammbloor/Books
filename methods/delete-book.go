package methods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DeleteBook() {
	bookBase, err := os.ReadFile("Book-base.txt")
	if err != nil { //
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	// вывод всех книг + скан номера нужной книги
	book := strings.Split(string(bookBase), "\n")

	fmt.Println("Выберите номер книги для удаления данных и введите в консоль!")
	for i, v := range book {
		fmt.Printf("%d: %s\n", i, v)
	}

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	bookNumber, _ := strconv.Atoi(scan.Text())

	if bookNumber < 0 || bookNumber > len(book)-1 {
		fmt.Println("Введите корректный номер")

	} else {
		tmpFile, err := os.Create("temp_file.txt")
		if err != nil {
			fmt.Println("Ошибка создания временного файла:", err)
			return
		}

		for i, line := range book {
			if i == bookNumber {
				continue
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
