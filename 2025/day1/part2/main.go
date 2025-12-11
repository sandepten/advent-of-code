package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	dialPointer := 50
	zeroCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Error converting to int: %v", err)
		}
		// add total number of hundreds
		zeroCount += hundreds(number)
		number = lastTwoDigits(number)

		if direction == "L" {
			currDialPointer := dialPointer - number
			if currDialPointer < 0 {
				if dialPointer != 0 && dialPointer != 100 {
					zeroCount++
				}
				dialPointer = 100 + currDialPointer
			} else {
				dialPointer = currDialPointer
			}
		} else {
			currDialPointer := dialPointer + number
			if currDialPointer > 100 {
				if dialPointer != 0 && dialPointer != 100 {
					zeroCount++
				}
				dialPointer = 0 + lastTwoDigits(currDialPointer)
			} else {
				dialPointer = currDialPointer
			}
		}
		if dialPointer == 0 || dialPointer == 100 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}

func lastTwoDigits(n int) int {
	res := n % 100
	if res < 0 {
		res += 100
	}
	return res
}

func hundreds(n int) int {
	return n / 100
}
