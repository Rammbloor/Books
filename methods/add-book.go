package methods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Books struct {
	name       string
	author     string
	sumOfPages int
}

// AddBook Эта функция открывает файл и записывает туда данные кооторые я передаю обращаясь к стуктуре

func (b Books) AddBook() {
	BookBase, err := os.OpenFile("Book-base.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil { // если возникла ошибка
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer BookBase.Close()
	fmt.Println("Введите название книги")
	name := bufio.NewScanner(os.Stdin)
	name.Scan()
	b.name = name.Text()
	_, err = BookBase.WriteString("|Book title| - " + b.name)
	if err != nil {
		return
	}
	fmt.Println("Введите Автора книги")
	author := bufio.NewScanner(os.Stdin)
	author.Scan()
	b.author = author.Text()
	_, err = BookBase.WriteString(" |Author| - " + b.author)
	if err != nil {
		return
	}
	fmt.Println("Введите количество страниц в книге")
	sumOfPages := bufio.NewScanner(os.Stdin)
	sumOfPages.Scan()
	b.sumOfPages, _ = strconv.Atoi(sumOfPages.Text())
	_, err = BookBase.WriteString(" |Number of pages| - " + strconv.Itoa(b.sumOfPages) + "\n")
	if err != nil {
		return
	}
}
