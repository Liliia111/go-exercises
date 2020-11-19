package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
	excTime := flag.Int64("time", 10, "time in s")
	flag.Parse()
	recordFile, err := os.Open(*fileName)

	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	reader := csv.NewReader(recordFile)
	records, _ := reader.ReadAll()
	correct := 0
	timer := time.NewTimer(time.Duration(*excTime) * time.Second)

problemloop:
	for _, r := range records {
		fmt.Printf("%v = ", r[0])
		textCh := make(chan string)
		go func() {
			var text string
			fmt.Scanln(&text)
			textCh <- text
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case text := <-textCh:
			if text == r[1] {
				correct++

			}
		}

	}
	fmt.Printf("You scored %v out of %v\n", correct, len(records))
}
