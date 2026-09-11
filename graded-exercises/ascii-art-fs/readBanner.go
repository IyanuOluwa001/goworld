package main

import (
	"os"
	"strings"
)

func readBannerFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	//Turns characters in my banner to a list
	lines := strings.Split(string(data), "\n")
	return lines, nil
}