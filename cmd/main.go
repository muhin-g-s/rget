package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/muhin-g-s/rget/internal/app"
)

func CountWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func main() {
	text := "Go is a simple and powerful language"
	count := CountWords(text)

	fmt.Println("Word count:", count)

	result := app.Run(os.Args[1:])

	fmt.Printf("%s", result)
}
