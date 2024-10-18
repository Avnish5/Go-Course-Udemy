package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Filemanager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm Filemanager) ReadLines() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read the file content")
	}

	return lines, nil

}

func (fm Filemanager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {

		return errors.New("Falied to create json file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {

		return errors.New("Falied to  convert data to json")
	}

	file.Close()
	return nil

}

func New(inputPath, outputPath string) Filemanager {
	return Filemanager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
