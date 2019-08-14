package main

import (
	"fmt"
	"log"
	"os"

	"github.com/naoty/ttodo/views"
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

	err := views.NewApplication().Run()

	if err != nil {
		log.Fatal(err)
	}
}
