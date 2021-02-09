package fileReader

import (
	"io/ioutil"
	"os"
	"errors"
)

func GetFileContent() (string, error) {
	if len(os.Args) < 1 {
		return "", errors.New("You must give a fileName in argument.")
	}

	fileName := os.Args[1]
	content, err := ioutil.ReadFile("./src/datasets/" + fileName)

	return string(content), err
}
