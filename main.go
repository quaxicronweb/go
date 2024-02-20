package main

import (
	"flag"
	"fmt"
)

func print(text string) {
	fmt.Println(text)
}

func main() {
	text := flag.String("text", "Hi", "text to print")
	flag.Parse()
	print(*text)
}
