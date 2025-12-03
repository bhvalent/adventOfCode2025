package day2

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bhvalent/adventOfCode2025/src/internal/utilities"
)

func GetInvalidRangeSum(filename string, includeAnySequenceLength bool) int {
	lines, err := utilities.GetLinesFromFile(filename)
	if err != nil || len(lines) != 1 {
		return 0
	}

	ranges := getRanges(lines[0])

	sum := 0
	for _, r := range ranges {
		sum += getRangeSum(r, includeAnySequenceLength)
	}

	return sum
}

func getRangeSum(r *Range, includeAnySequenceLength bool) int {
	sum := 0
	for i := r.start; i <= r.end; i++ {
		if (!includeAnySequenceLength && !isValidId(i)) || (includeAnySequenceLength && !isValidIdV2(i)) {
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

func isValidIdV2(num int) bool {
	snum := strconv.Itoa(num)

	mid := len(snum) / 2
  for seqLen := mid; seqLen > 0; seqLen-- {
    if len(snum) % seqLen != 0 {
      continue
    }

    if hasRepeatingSequence(snum, seqLen) {
      return false
    }
  }

  return true
}

func hasRepeatingSequence(data string, seqLen int) bool {
  seq := []string{}
  for i := 0; i < len(data); i += seqLen {
    if i + seqLen > len(data) {
      continue
    }

    seq = append(seq, data[i:i+seqLen])
  }

  for i := 1; i < len(seq); i++ {
    if seq[0] != seq[i] {
      return false
    } 
  }

  return true
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

