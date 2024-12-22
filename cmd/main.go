package main

import (
	"fmt"
	"os"

	"github.com/serikdev/go-tasks/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file> [output_file]")
		return
	}

	inputFile := os.Args[1]
	outputFile := "output.txt"

	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	producer := service.NewFileProducer(inputFile)
	presenter := service.NewFilePresenter(outputFile)

	svc := service.NewService(producer, presenter)

	err := svc.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Processing complete. Results written to:", outputFile)
	}
}
