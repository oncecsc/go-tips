package main

import (
	"embed"
	"fmt"
)

//go:embed dir1
var f embed.FS

//go:embed dir1
//go:embed dir2
var multiDirs embed.FS

func main() {
	// read singal file
	hellotxt, err := f.ReadFile("dir1/hello.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("dir1/hello.txt content is: %s\n", hellotxt)

	// read multi files

	dir1Contents, err := multiDirs.ReadFile("dir1/hello.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("dir1/hello.txt content is: %s\n", dir1Contents)
	dir2Contents, err := multiDirs.ReadFile("dir2/dir.txt")

	if err != nil {
		panic(err)
	}
	fmt.Printf("dir1/hello.txt content is: %s\n", dir2Contents)
}
