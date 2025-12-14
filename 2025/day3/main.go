package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Error closing file: %v", cerr)
		}
	}()

	scanner := bufio.NewScanner(file)

	totalJoltage := 0
	for scanner.Scan() {
		line := scanner.Text()
		prevHighest := 0
		maxJoltage := 0
		for v := range strings.SplitSeq(line, "") {
			currNumber, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Error converting to int: %v", err)
			}
			if prevHighest == 0 {
				prevHighest = currNumber
				continue
			}
			maxJoltage = max(maxJoltage, (prevHighest*10)+currNumber)
			if currNumber > prevHighest {
				prevHighest = currNumber
			}
		}
		totalJoltage += maxJoltage
	}
	fmt.Println(totalJoltage)
}
