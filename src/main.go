package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// create file
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// flags
	newItem := flag.String("n", "", "Add a new todo item")
	newItemNoTrunc := flag.String("new", "", "Add a new todo item")

	list := flag.Bool("l", true, "List todo items")
	listNoTrunc := flag.Blool("list", true, "List todo items")

	// removeItem := flag.String("r", "", "Remove a todo item")
	// parse
	flag.Parse()
	// args := flag.Args()

	// write new todos
	if *newItem != "" {
		_, err = f.WriteString(*newItem)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	if *newItemNoTrunc != "" {
		_, err = f.WriteString(*newItemNoTrunc)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	if *list == true || *listNoTrunc == true {
		f, err := os.Open("todo.txt")
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("File empty or does not exist")
				return
			}
			log.Fatal(err)
		}
		defer f.Close()

		fileScanner := bufio.NewScanner(f)

		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		f.Close()

		for _, line := range fileLines {
			fmt.Println(line)
		}
	}
}
