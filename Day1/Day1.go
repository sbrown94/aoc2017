package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	fmt.Println("Hello")
	input, err := readFromFileAsString("Day1.txt")
	if(err != nil) {
		fmt.Println("could not read file") 
	}
	store := -1
	captcha := 0
	first := 0
	for i, r := range input {
		if i == 0 {
			first, err = strconv.Atoi(string(r))
			if(err != nil) {
				fmt.Println("rune is not parsable") 
			}
		}
		d, err := strconv.Atoi(string(r))
		if(err != nil) {
			fmt.Println("rune is not parsable") 
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
}

func readFromFileAsString(fileName string) (string, error) {
	dat, err := ioutil.ReadFile(fileName)
	if(err != nil) {
		fmt.Println("could not read file")
		return "", err
	}
	return string(dat), nil
}