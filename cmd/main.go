package main

import (
	"fmt"
	"github.com/serikdev/go-tasks/internal/service"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file> [output_file]")
		return
	}

	inputFile := os.Args[1]
	outputFile := "output.txt" // значение по умолчанию

	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	// Создаем экземпляры продюсера и презентера
	producer := service.NewFileProducer(inputFile)
	presenter := service.NewFilePresenter(outputFile)

	// Создаем сервис
	svc := service.NewService(producer, presenter)

	// Запускаем сервис
	err := svc.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Processing complete. Results written to:", outputFile)
	}
}
