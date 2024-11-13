package main

func SpamMask(msg string) string {
	result := []byte{}
	buff := []byte(msg)
	length := len(buff)

	i := 0
	for i < length {
		if i+7 < length && buff[i] == 'h' && buff[i+1] == 't' && buff[i+2] == 't' && buff[i+3] == 'p' && buff[i+4] == ':' && buff[i+5] == '/' && buff[i+6] == '/' {
			result = append(result, buff[17:i+7]...)
			start := i + 7

			for i < length && buff[i] != ' ' {
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

func main() {
	msg := "Hidden web site: http://serdar.com"
	output := SpamMask(msg)
	println(output)
}
