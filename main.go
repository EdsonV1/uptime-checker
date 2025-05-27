package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"uptime-checker/checker"
	"uptime-checker/utils"
)

func main() {

	interval := flag.Duration("interval", 30*time.Second, "Interval between checks (used only if duration > 0)")
	duration := flag.Duration("duration", 0, "Total duration for repeated checks (e.g., 5m, 1h)")
	input := flag.String("i", "urls.txt", "Input path to list of URLs or file")
	output := flag.String("o", "check_results.log", "Output file for logs")

	flag.Parse()

	fmt.Println("🔍 Uptime Checker Starting...")
	fmt.Printf("📄 %s: %s\n🕒 Interval: %v\n⏳ Duration: %v\n📁 Output: %s\n\n",
		checker.CheckIdentityArgs(*input), *input, *interval, *duration, *output)

	if len(os.Args) > 1 {

		switch checker.CheckIdentityArgs(*input) {
		case checker.File:

			checker.CheckAllURL(*input, *output, *interval, *duration)

		case checker.URL:

			utils.RepeatCheck(*interval, *duration, func() {
				status := checker.CheckUrl(*input)
				utils.SaveResultToFile(*output, *input, status)
			})

		case checker.Directory:
			fmt.Println("Its directory please add the filename - ", *input)
		case checker.NonExistentLocalPath:
			fmt.Println("The file doesnt exist - ", *input)
		case checker.Unknown:
			fmt.Println("Something happen, try again! - ", *input)
		default:
			fmt.Println("Invalid PathType - ", *input)
		}

	} else {
		fmt.Println("Please provide a filename or URL as a command-line argument.")
		fmt.Println("Example: go run main.go my_document.txt")
		fmt.Println("Example: go run main.go https://example.com")
	}

}
