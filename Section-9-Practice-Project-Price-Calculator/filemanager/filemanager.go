package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {

	file, err := os.Open(path)

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

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)

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
