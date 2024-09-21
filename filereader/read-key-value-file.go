package filereader

import (
	"fmt"
	"os"
)

func ReadAllKeyValue() {
	HomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("failed to open file")
		return
	}

	filePath := fmt.Sprintf("%v/%v", HomeDir, "file-read-test.txt")
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buffer := make([]byte, 1)
	var content []byte
	for {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Printf("error %v \n", err)
		}

		if n == 0 {
			break
		}

		if buffer[0] == '\n' {
			break
		}

		content = append(content, buffer[0])
	}
}
