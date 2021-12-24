/*
Advent of Code 2020 day 17
Trick Shot
*/

package main

import (
	"aoc2021/shared"
	"fmt"
	"regexp"
	"strconv"
)

type Probe struct {
	x, y     int
	x_v, y_v int
}

// Find the trajectory with the highest arc
func SolvePart1(input []string) ([]int, int) {
	target := ParseInput(input[0])
	var peak, max_peak int
	var best_v []int
	var hit bool

	for x_v := 1; x_v < target[1]; x_v++ {
		probe := Probe{}
		hit, peak = false, 0

		// TODO: Nice try, but 45° is max range, not height
		for y_v := 0; y_v <= 10*x_v; y_v++ {
			hit, peak = GetArc(probe, target)

			if hit {
				if peak > max_peak {
					max_peak = peak
					best_v = []int{probe.x, probe.y}
				}
			}
		}
	}

	return best_v, max_peak
}

// ???
func SolvePart2(input []string) ([]int, int) {
	return []int{-1, -1}, -1
}

func ParseInput(input string) (out []int) {
	pattern, _ := regexp.Compile(`x=(\d+)\.\.(\d+), y=(-?\d+)\.\.(-?\d+)`)
	res := pattern.FindAllStringSubmatch(input, -1)

	for _, v := range res[0][1:] {
		v_i, _ := strconv.Atoi(v)
		out = append(out, v_i)
	}

	return
}

func GetArc(probe Probe, target []int) (hit bool, peak int) {
	var stop bool

	// Run until probe is below or right of target
	for !hit && !stop {
		hit, stop = Check(probe, target)

		if probe.y > peak {
			peak = probe.y
		}

		// Next physic step
		probe.x += probe.x_v
		probe.y += probe.y_v

		if probe.x_v > 0 {
			probe.x_v--
		} else if probe.x_v < 0 {
			probe.x_v++
		}

		probe.y_v--
	}

	return
}

func Check(probe Probe, target []int) (hit bool, stop bool) {
	// Probe lands within target bounds
	if probe.x >= target[0] && probe.x <= target[1] && probe.y >= target[2] && probe.y <= target[3] {
		return true, true
	}

	// Probe overshot target
	if probe.x > target[1] || probe.y < target[2] {
		return false, true
	}

	// Probe hasn't yet reached target
	return false, false
}

func main() {
	input := shared.ReadInputLines("day17/input")

	p1_v, p1_h := SolvePart1(input)
	p2_v, p2_h := SolvePart2(input)

	fmt.Printf("Part 1: %v: %v\n", p1_v, p1_h)
	fmt.Printf("Part 2: %v: %v\n", p2_v, p2_h)
}
