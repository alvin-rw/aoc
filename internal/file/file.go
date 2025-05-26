package file

import (
	"bufio"
	"log"
	"os"
)

// read file content and returns file lines as slice of string
// for AOC, failure to read the file content will always result in program exit
// no error returned to simplify function usage
func ReadFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error when opening the file: %v", err)
	}
	defer f.Close()

	fileContent := []string{}

	scanner := bufio.NewScanner(f)

	// increase buffer size to accomodate larger inputs
	bufferSize := 256 * 1024 // 256 KB
	buf := make([]byte, bufferSize)
	scanner.Buffer(buf, bufferSize)

	for scanner.Scan() {
		line := scanner.Text()
		fileContent = append(fileContent, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error when reading file: %v", err)
	}

	return fileContent
}
