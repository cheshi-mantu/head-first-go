package main

import (
	"fmt"
	"log"
	"os"
)

//fileinfo shows the size of fileinfo.go file

func main() {
	fileinfo, err := os.Stat("fileinfo.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file", fileinfo.Name(), "size is", fileinfo.Size())

}
