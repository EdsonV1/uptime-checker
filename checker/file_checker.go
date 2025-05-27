package checker

import (
	"fmt"
	"time"
	"uptime-checker/utils"
)

func CheckAllURL(file string, output string, interval time.Duration, duration time.Duration) {
	urlList, err := utils.ReadFromFile(file)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.RepeatCheck(interval, duration, func() {
		for _, url := range urlList {
			status := CheckUrl(url)
			utils.SaveResultToFile(output, file, status)
		}
	})
}
