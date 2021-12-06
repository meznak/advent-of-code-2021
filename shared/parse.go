package shared

import (
	"strconv"
	"strings"
)

func CSVToInt(input string) (output []int) {
	input_s := strings.Split(input, ",")

	for _, v := range input_s {
		v_i, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		output = append(output, v_i)
	}

	return
}
