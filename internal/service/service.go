package service

import "sync"

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present([]string) error
}

type Service struct {
	prod Producer
	pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}

func (s *Service) SpamMask(msg string) string {
	result := []byte{}
	buff := []byte(msg)
	leng := len(buff)

	i := 0
	for i < leng {
		if i+7 < leng &&
			buff[i] == 'h' &&
			buff[i+1] == 't' &&
			buff[i+2] == 't' &&
			buff[i+3] == 'p' &&
			buff[i+4] == ':' &&
			buff[i+5] == '/' &&
			buff[i+6] == '/' {
			result = append(result, buff[i:i+7]...)
			start := i + 7

			for i < leng && buff[i] != ' ' {
				i++
			}
			linkLen := i - start
			for j := 0; j < linkLen; j++ {
				result = append(result, '*')
			}

		} else {
			result = append(result, buff[i])
			i++
		}
	}
	return string(result)
}

func (s *Service) Run() error {
	lines, err := s.prod.Produce()
	if err != nil {
		return err
	}

	inputChan := make(chan string, len(lines))
	outputChan := make(chan string, len(lines))

	var wg sync.WaitGroup

	maxWorkers := 10
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range inputChan {
				outputChan <- s.SpamMask(line)
			}
		}()
	}

	go func() {
		for _, line := range lines {
			inputChan <- line
		}
		close(inputChan)
	}()

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	var maskedLines []string
	for maskedLine := range outputChan {
		maskedLines = append(maskedLines, maskedLine)
	}

	return s.pres.Present(maskedLines)
}
