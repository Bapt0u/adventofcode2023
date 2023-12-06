package helper

import (
	"log"
	"strconv"
	"strings"
)

func IntFromString(rawLine string, separator string) []int {
	line := strings.Split(rawLine, separator)
	ints := make([]int, len(line))
	for y, item := range line {
		item, err := strconv.Atoi(item)
		if err != nil {
			log.Println(err)
		}
		ints[y] = item
	}
	return ints
}
