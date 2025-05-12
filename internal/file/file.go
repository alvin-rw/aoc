package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error when opening the file: %w", err)
	}
	defer f.Close()

	fileContent := []string{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fileContent = append(fileContent, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error when reading file: %w", err)
	}

	return fileContent, nil
}
