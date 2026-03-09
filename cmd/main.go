package main

import (
	"fmt"
	"os"

	"github.com/muhin-g-s/rget/internal/app"
)

func main() {
	result := app.Run(os.Args[1:])

	fmt.Printf("%s", result)
}
