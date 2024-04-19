package modules

import (
	"bufio"
	"fmt"
	"os"
)

func CreateDatabase(filename string, name string) {
	// Check if the database file exists, if not, create it
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		fmt.Println("Created file named " + `"` + name + `"`)
		fmt.Println("")
		if err != nil {
			fmt.Println("Error creating database:", err)
			os.Exit(1)
		}
		defer file.Close()
	}
}

func DisplayNotes(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Notes:")
	for i := 1; scanner.Scan(); i++ {
		fmt.Printf("%03d - %s\n", i, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading database:", err)
	}
}

func AddNote(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the note text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error adding note:", err)
		return
	}

	fmt.Println("Note added successfully.")
}

func DeleteNote(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var notes []string
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	fmt.Print("Enter the number of note to remove or 0 to cancel: ")
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	if choice == 0 || choice > len(notes) {
		fmt.Println("Operation canceled or invalid note number.")
		return
	}

	notes = append(notes[:choice-1], notes[choice:]...)
	file.Truncate(0)
	file.Seek(0, 0)
	for _, note := range notes {
		file.WriteString(note + "\n")
	}

	fmt.Println("Note deleted successfully.")
}
