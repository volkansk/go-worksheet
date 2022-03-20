package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	openFile(filename)
}

func openFile(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error occurred while reading file", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
