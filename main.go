package main

import (
	"fmt"
	"os"

	"./pkg/stupidcd"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Wrong argument")
		os.Exit(1)
	}

	foundDirectory := stupidcd.PickOutDirectory(".", os.Args[1])

	if foundDirectory == "" {
		os.Exit(1)
	}
	fmt.Println(foundDirectory)
}
