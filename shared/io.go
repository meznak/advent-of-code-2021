package shared

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInputLines(fname string) []string {
	lines := make([]string, 0)

	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadInputInts(fname string) []int {
	lines := ReadInputLines(fname)
	ints := make([]int, len(lines))
	var err error

	for i, v := range lines {
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}

	return ints
}
