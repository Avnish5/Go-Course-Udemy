package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes-app/note"
)

func getNoteData() (string, string) {

	title := takeInput("Note Title: ")

	content := takeInput("Note Content: ")

	return title, content
}

func main() {

	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {

		fmt.Print(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {

		fmt.Print("Saving the note file failed")
		return
	}

	fmt.Print("Saving the note succedded")

}

func takeInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
