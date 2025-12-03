package day2

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bhvalent/adventOfCode2025/src/internal/utilities"
)

func GetInvalidRangeSum(filename string) int {
	lines, err := utilities.GetLinesFromFile(filename)
	if err != nil || len(lines) != 1 {
		return 0
	}

	ranges := getRanges(lines[0])

	sum := 0
	for _, r := range ranges {
		sum += getRangeSum(r)
	}

	return sum
}

func getRangeSum(r *Range) int {
	sum := 0
	for i := r.start; i <= r.end; i++ {
		if !isValidId(i) {
			sum += i
		}
	}
	return sum
}

func isValidId(num int) bool {
	snum := strconv.Itoa(num)

	if len(snum)%2 == 1 {
		return true
	}

	mid := len(snum) / 2
	return !(snum[:mid] == snum[mid:])
}

type Range struct {
	start int
	end   int
}

func getRanges(data string) []*Range {
	ranges := []*Range{}
	dataPoints := strings.Split(data, ",")
	for _, dp := range dataPoints {
		r, err := newRange(dp)
		if err != nil {
			continue
		}
		ranges = append(ranges, r)
	}
	return ranges
}

func newRange(dp string) (*Range, error) {
	ends := strings.Split(dp, "-")

	if len(ends) != 2 {
		return nil, errors.New("Missing end of range")
	}

	s, sErr := strconv.Atoi(ends[0])
	e, eErr := strconv.Atoi(ends[1])

	if sErr != nil || eErr != nil {
		return nil, errors.New("Error converting to int")
	}

	r := Range{
		start: s,
		end:   e,
	}

	return &r, nil
}

