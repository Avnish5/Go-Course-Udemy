package main

import (
	"bufio"
	"example.com/interface-app/note"
	"example.com/interface-app/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func getNoteData() (string, string) {

	title := takeInput("Note Title: ")

	content := takeInput("Note Content: ")

	return title, content
}

func getTodoData() string {

	text := takeInput("Enter your todo: ")

	return text
}

func main() {

	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		return
	}

	err = outputdata(userNote)

	if err != nil {
		return
	}

	fmt.Print("Saving the note succedded")

	todoText := getTodoData()
	todo, err := todo.New(todoText)

	if err != nil {

		fmt.Print(err)
		return
	}

	err = outputdata(todo)

	if err != nil {
		return
	}

	fmt.Print("Saving the todo succedded")

}

func outputdata(data outputtable) error {
	data.Display()

	return saveData(data)
}
func saveData(data saver) error {
	err := data.Save()

	if err != nil {

		fmt.Print("Saving the  file failed")
		return err
	}

	return nil
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
