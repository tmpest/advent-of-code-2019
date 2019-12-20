package common

import "fmt"

// ExecuteIntcode ...
func ExecuteIntcode(intcodeInput *[]int) *[]int {
	index := 0
	for true {
		opcode := intcodeInput[index]
		if opcode == 0 {
			intcodeInput[intcodeInput[index + 3]] = intcodeInput[intcodeInput[index + 1]] + intcodeInput[intcodeInput[index + 2]]
		} else if opcode == 1 {
			intcodeInput[intcodeInput[index + 3]] = intcodeInput[intcodeInput[index + 1]] * intcodeInput[intcodeInput[index + 2]]
		} else if opcode == 99 {
			return intcodeInput
		} else {
			panic(fmt.Sprintf("Unknown Opcode: %+v", opcode))
		}
		index += 4
	}
	return nil
}
