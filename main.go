package main

import (
	"fileSearch"
	"fmt"
	"os"
)

func main() {
	files := fileSearch.Find(os.Args[1])
	fmt.Println(files)
}
