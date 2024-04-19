package modules

import (
	"bufio"
	"fmt"
	"os"
)

const dataFolder = "./data/"

func Manager() {
	fmt.Println("Welcome to the Data Manager!")

	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show files.")
		fmt.Println("2. Create a file.")
		fmt.Println("3. Delete a file.")
		fmt.Println("4. Exit.")

		var choice int
		fmt.Print(">")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			showFiles()
		case 2:
			createFile()
		case 3:
			deleteFile()
		case 4:
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func showFiles() {
	files, err := os.ReadDir(dataFolder)
	if err != nil {
		fmt.Println("Error reading data folder:", err)
		return
	}

	fmt.Println("Files in the data folder:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func createFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the file to create: ")
	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fileName = dataFolder + fileName[:len(fileName)-1] // Remove newline character

	if _, err := os.Stat(fileName); err == nil {
		fmt.Println("File already exists.")
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully.")
}

func deleteFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the file to delete: ")
	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fileName = dataFolder + fileName[:len(fileName)-1] // Remove newline character

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("File does not exist.")
		return
	}

	err = os.Remove(fileName)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully.")
}
