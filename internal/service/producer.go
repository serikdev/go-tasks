package service

import (
	"bufio"
	"os"
)

type FileProduce struct {
	filePath string
}

func NewFileProducer(filePath string) *FileProduce {
	return &FileProduce{filePath: filePath}
}

func (fpd *FileProduce) Produce() ([]string, error) {
	file, err := os.Open(fpd.filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
