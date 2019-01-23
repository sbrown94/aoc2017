package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sbrown94/aoc2017/common"
)

type Day struct {
}

func New() Day {
	return Day{}
}

// RunP1 Runs the first part of the challenge
func (d Day) RunP1() error {
	input, err := common.ReadFromFileAsString("./day2/day2.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		winHigh := 0
		winLow := 9999999
		nums := strings.Split(line, "\t")
		fmt.Println(nums)
		for _, n := range nums {
			nI, err := strconv.Atoi(n)
			if err != nil {
				return errors.WithStack(err)
			}
			if nI > winHigh {
				winHigh = nI
			}
			if nI < winLow {
				winLow = nI
			}
		}
		fmt.Println(winHigh - winLow)
		total += (winHigh - winLow)
	}
	fmt.Println(total)
	return nil
}

// RunP2 Runs the second part of the challenge
func (d Day) RunP2() error {
	return nil
}
