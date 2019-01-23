package day1

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sbrown94/aoc2017/common"
)

type Day struct {
}

func New() Day {
	return Day{}
}

// RunP1 Part 1 of the problem
func (d Day) RunP1() error {
	input, err := common.ReadFromFileAsString("./day1/day1.txt")
	if err != nil {
		return err
	}
	store := -1
	captcha := 0
	first := 0
	for i, r := range input {
		if i == 0 {
			first, err = strconv.Atoi(string(r))
			if err != nil {
				return errors.WithStack(err)
			}
		}
		d, err := strconv.Atoi(string(r))
		if err != nil {
			return errors.WithStack(err)
		}
		if d == store {
			captcha += d
		}
		store = d
	}
	if store == first {
		captcha += first
	}
	fmt.Println(captcha)
	return nil
}

// RunP2 Part 2 of the problem
func (d Day) RunP2() error {
	input, err := common.ReadFromFileAsString("./day1/day1.txt")
	if err != nil {
		return err
	}
	numList, err := common.StringToIntSlice(input)
	if err != nil {
		return err
	}
	jumper := len(numList) / 2
	total := 0
	for i, num := range numList {
		if i+jumper >= len(numList) {
			jumper -= len(numList)
		}
		if num == numList[i+jumper] {
			total += num
		}
	}
	fmt.Println(total)
	return nil
}
