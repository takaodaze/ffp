package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ffp <target_file>")
		os.Exit(1)
	}

	file := os.Args[1]
	wd, getwdErr := os.Getwd()

	if getwdErr != nil {
		fmt.Fprintln(os.Stderr, getwdErr)
		os.Exit(1)
	}

	var fileFullpath string = wd + "/" + file

	if _, err := os.Stat(fileFullpath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", fileFullpath)
}
