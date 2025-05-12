package file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error when opening the file, %v", err)
	}
	defer f.Close()

	fileContent := []string{}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, fmt.Errorf("error when reading file, %v", err)
			}
		}
		line = strings.TrimSuffix(line, "\n")

		fileContent = append(fileContent, line)
	}

	return fileContent, nil
}
