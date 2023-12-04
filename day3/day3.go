package day3

import (
	"bufio"
	"log"
	"os"
)

func GetAdjacentNumbers(input string) {
	matrix := LoadMatrix(input)
	var values []int
	var previousWasChar bool = false
	var pow int = 1
	log.Println(len(matrix))
	for x, line := range matrix {
		for y, char := range line {
			if char >= 48 && char <= 57 {

			}
		}
	}
}

func LoadMatrix(file string) []string {
	var matrix []string
	log.Println("Opening", file)

	readFile, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		matrix = append(matrix, fileScanner.Text())
	}

	readFile.Close()

	return matrix
}
