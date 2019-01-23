package common

import (
	"io/ioutil"
	"strconv"
)

// ReadFromFileAsString takes a filename and returns the contents as a string
func ReadFromFileAsString(fileName string) (string, error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

// StringToIntSlice takes a string and converts it into a slice of ints of equivalent length
func StringToIntSlice(str string) ([]int, error) {
	var slc []int
	for _, c := range str {
		cI, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}
		slc = append(slc, cI)
	}
	return slc, nil
}
