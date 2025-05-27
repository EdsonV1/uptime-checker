package utils

import (
	"fmt"
	"os"
	"time"
)

func SaveResultToFile(filePath, url, status string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("%s | %s | %s\n", timestamp, url, status)
	_, err = f.WriteString(entry)
	return err
}
