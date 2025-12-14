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
	var dataLine string
	for scanner.Scan() {
		dataLine = scanner.Text()
	}

	for v := range strings.SplitSeq(dataLine, ",") {
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
			for dividePoint := 1; dividePoint <= numLength/2; dividePoint++ {
				if numLength%dividePoint != 0 {
					continue
				}
				allMatch := true
				matchPart := ""
				for j := 0; j < numLength/dividePoint; j++ {
					divideAt := j * dividePoint
					currPart := strconv.Itoa(i)[divideAt:][:dividePoint]
					if len(matchPart) == 0 {
						matchPart = currPart
					} else if matchPart != currPart {
						allMatch = false
						break
					}
				}
				if allMatch {
					totalSum += i
					break
				}
			}
		}
	}

	fmt.Println(totalSum)
}
