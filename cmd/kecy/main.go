package main

import (
	"fmt"
	"os"

	"github.com/jiikko/Karabiner-Elements-config-yaml/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <yaml filepath>")
		return
	}

	filepath := os.Args[1]
	parser, err := internal.NewParser(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(parser.ToJSON())
}
