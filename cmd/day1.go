package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tmpest/advent-of-code-2019/pkg/common"
)

func main() {
	fmt.Printf("Total Fuel Required...for real: %+v", partTwo())
}

func partOne() int {
	path, e := filepath.Abs("resources/input/1.0.txt")
	check(e)

	fileInput, e := ioutil.ReadFile(path)
	check(e)

	inputs := strings.Split(string(fileInput), "\n")
	sum := 0
	for _, inputString := range inputs {
		inputString = strings.Trim(inputString, "\r")
		input, e := strconv.ParseInt(inputString, 10, 0)
		check(e)
		sum += common.CalculateRequiredFuel(int(input))
	}
	return sum
}

func partTwo() int {
	path, e := filepath.Abs("resources/input/1.0.txt")
	check(e)

	fileInput, e := ioutil.ReadFile(path)
	check(e)
	inputs := strings.Split(string(fileInput), "\r\n")
	sum := 0
	for _, inputString := range inputs {
		// inputString = strings.Trim(inputString, "\r")
		input, e := strconv.ParseInt(inputString, 10, 0)
		check(e)

		sum += common.CalculateRequiredFuelIncludingFuelMass(int(input))
	}
	return sum
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
