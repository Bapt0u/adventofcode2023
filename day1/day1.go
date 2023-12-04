package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetCalibrationValue(line string) int {
	var firstNum, lastNum int
	var firstValFound bool = false
	for _, y := range line {
		if val, err := strconv.Atoi(string(y)); err == nil {
			if !firstValFound {
				firstValFound = true
				firstNum = val
			} else {
				lastNum = val
			}
		}
	}
	if lastNum == 0 {
		lastNum = firstNum
	}
	return (firstNum * 10) + lastNum
}

func ComputeFileForElves(file string) int {
	readFile, err := os.Open(file)
	var total int = 0
	if err != nil {
		log.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		total += GetCalibrationValue(fileScanner.Text())
	}

	readFile.Close()

	return total
}
