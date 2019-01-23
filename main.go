package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sbrown94/aoc2017/day1"
)

type day interface {
	RunP1() error
	RunP2() error
}

var days = []day{
	day1.New(),
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var dayNo = flag.Int("day", 0, "Please enter a day to run")
	var part = flag.Bool("p2", false, "Run the second part of the challenge")
	flag.Parse()

	if *dayNo == 0 {
		log.Fatal("Day flag must be set")
	}

	day, err := getDay(*dayNo)
	if err != nil {
		log.Fatal(err)
	}

	err = executePart(day, *part)
	if err != nil {
		log.Fatal(err)
	}

}

func getDay(dayNo int) (day, error) {
	if dayNo > len(days) {
		return nil, fmt.Errorf("Day '%d' is not available", dayNo)
	}
	return days[dayNo-1], nil
}

func executePart(day day, p2 bool) error {
	if p2 {
		return day.RunP2()
	}
	return day.RunP1()
}
