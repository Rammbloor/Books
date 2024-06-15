package methods

import (
	"fmt"
	"os"
	"strings"
)

func ShowAll() {

	BookBase, err := os.ReadFile("Book-base.txt")
	{
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	Book := strings.Split(string(BookBase), "\n")

	for i, v := range Book {
		fmt.Printf("%d: %s\n", i, v)

	}
}
