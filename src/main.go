package main

import (
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

	removeItem := flag.String("r", "", "Remove a todo item")
	// parse
	flag.Parse()
	// args := flag.Args()

	// write new todos
	_, err = f.WriteString(*newItem)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	// fmt.Println("Args:", args)
	fmt.Println("Todo Item:", *newItem)
	fmt.Println("Todo Item:", *newItemNoTrunc)
	fmt.Println("Listing items: ", *list)
	fmt.Println("Removing item: ", *removeItem)
}
