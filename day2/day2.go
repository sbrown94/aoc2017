package day2

import (
	"fmt"
	"math"
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
		total += (winHigh - winLow)
	}
	fmt.Println(total)
	return nil
}

// RunP2 Runs the second part of the challenge
func (d Day) RunP2() error {
	input, err := common.ReadFromFileAsString("./day2/day2.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(input, "\n")
	var total float64
	for _, line := range lines {
		nums := strings.Split(line, "\t")
		for _, outerNum := range nums {
			for _, innerNum := range nums {
				if innerNum == outerNum {
					continue
				}
				outerNumFloat, err := strconv.ParseFloat(outerNum, 64)
				if err != nil {
					return errors.WithStack(err)
				}
				innerNumFloat, err := strconv.ParseFloat(innerNum, 64)
				if err != nil {
					return errors.WithStack(err)
				}
				//fmt.Println(math.Mod(outerNumFloat, innerNumFloat))
				if math.Mod(outerNumFloat, innerNumFloat) == float64(0) {
					total += (outerNumFloat / innerNumFloat)
					break
				}
			}
		}
	}
	fmt.Println(total)
	return nil
}
