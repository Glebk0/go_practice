package main

import (
	"fmt"
	"io"
	"os"
)

// Read the contents of a text file, then print its contents to the terminal

func main() {
	f_content, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f_content)

}
