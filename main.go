package main

import (
	"fmt"
	"os"
)

// Version represents the version of this application.
var Version string

func main() {
	for _, arg := range os.Args {
		switch arg {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		}
	}
}
