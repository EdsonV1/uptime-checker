package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFromFile(filePath string) ([]string, error) {

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	listUrl := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			listUrl = append(listUrl, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
		return nil, err
	}

	return listUrl, nil
}
