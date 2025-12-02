package day1

import (
	"fmt"
	"strconv"

	"github.com/bhvalent/adventOfCode2025/src/internal/utilities"
)

func HowManyTimesAtZero(filename string) int {
	lines, err := utilities.GetLinesFromFile(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	// get instructions
	instructions := getInstructions(lines)

	// count times at zero
	timesAtZero := 0
	startingPosition := 50
	currentPosition := startingPosition
	for _, i := range instructions {
		distance := i.distance
		if i.direction == "L" {
			distance = distance * -1
		}
		currentPosition = getUpdatedPosition(currentPosition, distance)
		fmt.Printf("Position: %d\n", currentPosition)
		if currentPosition == 0 {
			timesAtZero += 1
		}
	}

	return timesAtZero
}

func getUpdatedPosition(current int, distance int) int {
	value := current + distance
	value = value % 100
	if value < 0 {
		value = value + 100
	}
	return value
}

type Instruction struct {
	direction string
	distance  int
}

func getInstructions(lines []string) []*Instruction {
	if len(lines) < 1 {
		return nil
	}

	instructions := []*Instruction{}
	for _, value := range lines {
		if len(value) < 1 {
			continue
		}

		instructions = append(instructions, newInstruction(value))
	}

	return instructions
}

func newInstruction(data string) *Instruction {
	i := Instruction{}
	if len(data) < 2 {
		return &i
	}

	i.direction = string(data[0])
	distance, err := strconv.Atoi(data[1:])
	if err != nil {
		distance = 0
	}
	i.distance = distance
	return &i
}

