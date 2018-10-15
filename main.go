package main

import (
	"fmt"
	"os"
)

func main() {
	files := Find(os.Args[1])
	fmt.Println(files)
}
