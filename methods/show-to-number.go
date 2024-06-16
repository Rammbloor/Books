package methods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowBookToNumber() {
	fmt.Println("Введите номер книги")

	bookBase, err := os.ReadFile("Book-base.txt")
	{
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	book := strings.Split(string(bookBase), "\n")
	for {
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		numberBook, _ := strconv.Atoi(sc.Text())
		if numberBook < 0 || numberBook > len(book)-1 {
			fmt.Println("Введите корректный номер")
			continue
		}
		for i, v := range book {

			if i == numberBook {
				fmt.Printf("%d: %s\n", i, v)
				os.Exit(0)
			}
		}
	}
}
