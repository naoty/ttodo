package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/naoty/ttodo/todo"
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

	file, err := os.Open(todoPath())

	if err != nil {
		log.Fatal(err)
	}

	store := todo.GetStore(file)
	err = store.LoadTodos()

	if err != nil {
		log.Fatal(err)
	}

	err = views.NewApplication().Run()

	if err != nil {
		log.Fatal(err)
	}
}

func todoPath() string {
	dir := os.Getenv("TODO_PATH")
	if dir == "" {
		dir = os.Getenv("HOME")
	}

	return filepath.Join(dir, ".todo.json")
}
