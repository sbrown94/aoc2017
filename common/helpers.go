package common

import (
	"io/ioutil"
)

// ReadFromFileAsString takes a filename and returns the contents as a string
func ReadFromFileAsString(fileName string) (string, error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
