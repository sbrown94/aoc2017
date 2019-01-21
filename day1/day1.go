package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type Day struct {
	captcha int
}

func New() Day {
	captcha := RunP1()
	return Day{
		captcha: captcha,
	}
}

func RunP1() int {
	input, err := readFromFileAsString("day1.txt")
	if err != nil {
		log.Fatal("could not read file")
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
	return captcha
}

func readFromFileAsString(fileName string) (string, error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("could not read file")
		return "", err
	}
	return string(dat), nil
}
