package main

import (
	"io/ioutil"
	"os"
)

func OpenFile(path string) ([]byte, error) {

	// open the file
	// create if not exist
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	// read the content using buffer
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
