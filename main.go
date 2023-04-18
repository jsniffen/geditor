package main

import "fmt"
import "github.com/jsniffen/geditor/term"
import "log"

func main() {
	err := term.Init()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 255; i += 10 {
		term.PrintColor(uint8(i), uint8(i), uint8(i))
		fmt.Println("hello world")
	}
}
