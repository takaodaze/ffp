package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ffp <target_file>")
		os.Exit(1)
	}

	file := os.Args[1]
	wd, getwdErr := os.Getwd()

	if getwdErr != nil {
		fmt.Println(getwdErr)
		os.Exit(1)
	}

	var fileFullpath string = wd + "/" + file

	if _, err := os.Stat(fileFullpath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", fileFullpath)
}
