package utilities

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func GetLinesFromFile(str string) ([]string, error) {
	lines := []string{}

	if str == "" {
		return lines, errors.New("filename cannot be empty")
	}

	file, err := os.Open(str)

	if err != nil {
		errMsg := fmt.Sprintf("error opening %s: %s", str, err)
		return lines, errors.New(errMsg)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
