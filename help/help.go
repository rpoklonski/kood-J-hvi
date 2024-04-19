package help

import "fmt"

func Help() {
	fmt.Print(`notestool creates and manages your notes.

Usage:

		go run main.go <filename> 

filename can be one-line, words can be separated with underscore "_"
only one filename can be used

In notestool you are shown what options you can take and you can choose them
by corresponding number option.

Creating note will ask you to write a string (text) of a note you'd like to save

Deleting note will ask you wich note youd like to delet by position number

`)
}
