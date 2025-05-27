package checker

import (
	"fmt"
	"net/http"
	"time"
)

func CheckUrl(url string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("❌ %s is DOWN. Error: %v\n", url, err)
		return fmt.Sprintf("DOWN: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Printf("✅ %s is UP. Status: %s\n", url, resp.Status)
		return "UP"
	} else {
		fmt.Printf("⚠️ %s returned status: %s\n", url, resp.Status)
		return fmt.Sprintf("WARNING: %s", resp.Status)
	}
}
