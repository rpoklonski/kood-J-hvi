package main

import (
	"fmt"
	"os"
	"package/help"
	"package/modules"
)

const dataFolder = "./data/"

func main() {
	// Check if an argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>\nUsage: go run main.go --help")
		return
	}

	if os.Args[1] == "--help" {
		help.Help()
		return
	}
	if os.Args[1] == "-manager" {
		modules.Manager()
		return
	}

	// Get collection name from command line argument
	collection := os.Args[1]
	databaseFile := dataFolder + collection

	// Create or load the database
	modules.CreateDatabase(databaseFile, collection)

	// Display welcome message
	fmt.Println("Welcome to the notes tool!")

	// Menu loop
	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		// Read user input
		var choice int
		fmt.Println("")
		fmt.Print(">")
		_, err := fmt.Scanf("%d", &choice)
		fmt.Println("")
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			modules.DisplayNotes(databaseFile)
		case 2:
			modules.AddNote(databaseFile)
		case 3:
			modules.DeleteNote(databaseFile)
		case 4:
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}
