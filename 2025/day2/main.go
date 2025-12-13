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

	totalSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for v := range strings.SplitSeq(line, ",") {
			start, err := strconv.Atoi(strings.Split(v, "-")[0])
			if err != nil {
				log.Fatalf("Error converting to int: %v", err)
			}
			end, err := strconv.Atoi(strings.Split(v, "-")[1])
			if err != nil {
				log.Fatalf("Error converting to int: %v", err)
			}

			for i := start; i <= end; i++ {
				numLength := len(strconv.Itoa(i))
				if numLength%2 == 0 {
					mid := numLength / 2
					left := strconv.Itoa(i)[mid:]
					right := strconv.Itoa(i)[:mid]
					if left == right {
						totalSum += i
					}
				}
			}
		}
	}
	fmt.Println(totalSum)
}
