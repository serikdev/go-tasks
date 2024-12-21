package service

import (
	"bufio"
	"os"
)

type FilePresent struct {
	filePath string
}

func NewFilePresenter(filePath string) *FilePresent {
	return &FilePresent{filePath: filePath}
}

func (fps *FilePresent) Present(lines []string) error {
	file, err := os.Create(fps.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}
