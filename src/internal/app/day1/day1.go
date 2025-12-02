package day1

import (
	"fmt"
	"math"
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
		if currentPosition == 0 {
			timesAtZero += 1
		}
	}

	return timesAtZero
}

func HowManyTimesPassingZero(filename string) int {
	lines, err := utilities.GetLinesFromFile(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	// get instructions
	instructions := getInstructions(lines)

	// count times at zero
	timesPassingZero := 0
	startingPosition := 50
	currentPosition := startingPosition
	for _, i := range instructions {
		distance := i.distance
		if i.direction == "L" {
			distance = distance * -1
		}
		updatedPosition, passes := getUpdatedPositionAndPasses(currentPosition, distance)
		currentPosition = updatedPosition
		timesPassingZero += passes
	}

	return timesPassingZero
}

func getUpdatedPosition(current int, distance int) int {
	value := current + distance
	value = value % 100
	if value < 0 {
		value += 100
	}
	return value
}

func getUpdatedPositionAndPasses(current int, distance int) (int, int) {
	value := current + distance

	// calculate update value
	newValue := value % 100
	if newValue < 0 {
		newValue += 100
	}

	passes := 0
	// full rotations
	if distance <= -100 || distance >= 100 {
		passes += int(math.Abs(float64(distance))) / 100
	}

	// partial rotations
	isRightPartialPass := distance > 0 && newValue < current
	isLeftPartialPass := distance < 0 && (newValue > current && current != 0 || (newValue == 0 && current != 0))
	if isRightPartialPass || isLeftPartialPass {
		passes += 1
	}

	return newValue, passes
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

