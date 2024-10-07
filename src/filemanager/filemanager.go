package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

const basePath = "./store/"

func New(inputPath, outputPathPath string) FileManager {
	return FileManager{
		InputFilePath:  basePath + inputPath,
		OutputFilePath: basePath + outputPathPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		file.Close()
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, err
	}

	file.Close()

	return lines, nil
}

func (fm FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		file.Close()
		return errors.New("failed to create a file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to Json")
	}
	file.Close()
	return nil
}
