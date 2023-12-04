package day2

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsThisGamePossible(game string) int {
	var counter = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	var limitCounter = map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	gameInfo := strings.Split(game, ":")
	gameId, err := strconv.Atoi(regexp.MustCompile("[0-9]+").FindAllString(gameInfo[0], 1)[0])
	if err != nil {
		log.Println(err)
	}
	log.Printf("Id is %d", gameId)
	for _, sets := range strings.Split(gameInfo[1], ";") {
		for _, set := range strings.Split(sets, ",") {
			subset := strings.Split(set, " ")
			subsetInt, _ := strconv.Atoi(string(subset[1]))
			counter[string(subset[2])] += subsetInt

			if subsetInt > limitCounter[string(subset[2])] {
				log.Println(counter[string(subset[2])], string(subset[2]), " balls. exiting game.")
				return 0
			}
		}
	}
	return gameId
}

func IsThisGamePossible2(game string) int {
	var counter = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	gameInfo := strings.Split(game, ":")
	gameId, err := strconv.Atoi(regexp.MustCompile("[0-9]+").FindAllString(gameInfo[0], 1)[0])
	if err != nil {
		log.Println(err)
	}
	log.Printf("Id is %d", gameId)
	for _, sets := range strings.Split(gameInfo[1], ";") {
		for _, set := range strings.Split(sets, ",") {
			subset := strings.Split(set, " ")
			subsetInt, _ := strconv.Atoi(string(subset[1]))
			// counter[string(subset[2])] += subsetInt

			if subsetInt > counter[string(subset[2])] {
				counter[string(subset[2])] = subsetInt
			}
		}
	}
	return counter["green"] * counter["blue"] * counter["red"]
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
		total += IsThisGamePossible2(fileScanner.Text())
	}

	readFile.Close()

	return total
}
