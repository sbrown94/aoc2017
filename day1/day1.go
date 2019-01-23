package day1

import (
	"fmt"
	"log"
	"strconv"

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
		log.Fatal(err)
	}
	store := -1
	captcha := 0
	first := 0
	for i, r := range input {
		if i == 0 {
			first, err = strconv.Atoi(string(r))
			if err != nil {
				log.Fatal("rune is not parsable")
			}
		}
		d, err := strconv.Atoi(string(r))
		if err != nil {
			log.Fatal("rune is not parsable")
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
		log.Fatal(err)
	}
	doubleInput := input + input
	total := 0
	for i, r := range input {
		val, err := strconv.Atoi(string(doubleInput[(i + (len(input) / 2))]))
		if err != nil {
			return err
		}
		rangeVal, err := strconv.Atoi(string(r))
		if err != nil {
			return err
		}
		if rangeVal == val {
			total += val
		}
	}

	fmt.Println(total)

	return nil
}
