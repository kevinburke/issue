package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/browser"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: issue <number>\n")
		os.Exit(2)
	}
	err := browser.OpenURL("https://github.com/golang/go/issues/" + flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(2)
	}
}
