package common

import "fmt"

// ExecuteIntcode ...
func ExecuteIntcode(intcodeInput []int) []int {
	index := 0
	for index <= len(intcodeInput) {
		switch opcode := intcodeInput[index]; opcode {
		case 0:
			intcodeInput[intcodeInput[index+3]] = intcodeInput[intcodeInput[index+1]] + intcodeInput[intcodeInput[index+2]]
		case 1:
			intcodeInput[intcodeInput[index+3]] = intcodeInput[intcodeInput[index+1]] * intcodeInput[intcodeInput[index+2]]
		case 99:
			return intcodeInput
		default:
			panic(fmt.Sprintf("Unknown Opcode: %+v", opcode))
		}
		index += 4
	}
}
